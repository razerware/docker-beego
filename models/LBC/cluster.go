package LBC

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
	"time"
)

type ClusterStat struct {
	SwarmID   string `json:"swarm_id"`
	Token     string `json:"token"`
	ManagerIP string `json:"manager_ip"`
	ELBFlag   int
	Ready     int `json:"ready"`
	models.User
	models.ElasticInfo
}

func CheckCluster(record map[string]interface{}) ClusterStat {
	var cs ClusterStat
	//map->json->struct
	//glog.Info("transfer map to json")

	if err := MapToStruct(record, &cs); err != nil {
		glog.Error(err)
	} else {
		glog.V(1).Info("map transfer ok")
	}
	glog.Info("check cluster ",cs.SwarmID)
	//check if cluster is resizing
	if cs.Ready == 0 {
		cs.ELBFlag = 0
		glog.Error("This cluster is resizing")
		return cs
	}
	kv:=map[string]string{"swarm_id":cs.SwarmID,"role":"worker"}
	CpuSql := InfluxSelect("mean", "cpu", cs.Username, "autogen", "vm",
		"30s", kv)
	MemSql := InfluxSelect("mean", "mem", cs.Username, "autogen", "vm",
		"30s", kv)
	ir := PullData(cs.Username, CpuSql+";"+MemSql)
	glog.Info("..........................Judging Cluster: ",cs.SwarmID,".............................")
	cs.ELBFlag= Judgement(ir, cs)
	if cs.ELBFlag == 0 {
		return cs
	}
	amount := clusterMap.Get(cs.SwarmID)
	glog.Info("ELBflag != 0 ",cs.ELBFlag,"cluster's node is :" ,amount)
	return cs

}
func CountClusterNode() {
	sql := "SELECT COUNT(*)AS amount,swarm_id FROM `vm_info` WHERE swarm_id!='' GROUP BY swarm_id"
	record := models.MysqlQuery(sql)
	if record != nil {
		for _, i := range record {
			glog.Info("Counting node in swarm:",i["swarm_id"])
			clusterMap.Set(i["swarm_id"], i["amount"])
		}
		glog.V(1).Info(clusterMap)
	}
}

//checkExpand get nodeList
//set 0 for cluster ready
//join cluster
//update service sort by cpu usage desc
//set 1 to cluster ready
func (cs *ClusterStat) ExpandCluster() error {
	//nodeList is ip array
	nodeIDList, ok := cs.CheckExpand()
	if !ok {
		glog.Error("CheckExpand error")
		return fmt.Errorf("CheckExpand false")
	}
	//set 0 for cluster when adjusting cluster
	sql := fmt.Sprintf("UPDATE `cluster_info` SET `ready` = '0' WHERE `cluster_info`.`swarm_id` = '%s';",
		cs.SwarmID)
	models.MysqlUpdate(sql)

	defer func() {
		sql = fmt.Sprintf("UPDATE `cluster_info` SET `ready` = '1' WHERE `cluster_info`.`swarm_id` = '%s';",
			cs.SwarmID)
		_, _, err := models.MysqlUpdate(sql)
		if err != nil {
			glog.Error("ready Cluster error", err)
		}
	}()

	nodeIDList=nodeIDList[:cs.Step]
	ok = cs.JoinCluster(nodeIDList)
	if !ok {
		return fmt.Errorf("node join error")
	}
	ok=cs.UpdateCluster()
	if !ok {
		return fmt.Errorf("Update cluster error")
	}
	//update services
	influxsql := fmt.Sprintf("SELECT mean(cpu) AS mean_cpu FROM %s.autogen.container " +
		"WHERE time > now() - 5m GROUP BY service_id",
		cs.Username)
	ir := PullData("lzy", influxsql)
	serviceList := ir.SortInfluxResp("desc")
	err:=cs.UpdateServices(serviceList)
	return err
}

//set cluster to unready
//drain node y sort by cpu usage
//wait for container restart on other nodes
//leave node form cluster
func (cs ClusterStat) ReduceCluster() error {
	if !cs.CheckReduce() {
		glog.Error("checkReduce error")
		return fmt.Errorf("checkReduce error")
	}
	sql := fmt.Sprintf("UPDATE `cluster_info` SET `ready` = '0' WHERE `cluster_info`.`swarm_id` = '%s';",
		cs.SwarmID)
	models.MysqlUpdate(sql)

	defer func() {
		sql = fmt.Sprintf("UPDATE `cluster_info` SET `ready` = '1' WHERE `cluster_info`.`swarm_id` = '%s';",
			cs.SwarmID)
		_, _, err := models.MysqlUpdate(sql)
		if err != nil {
			glog.Error("ready Cluster error", err)
		}
	}()

	influxsql := fmt.Sprintf("SELECT mean(cpu) AS mean_cpu FROM %s.autogen.vm "+
		"WHERE time > now() - 5m AND role='worker' GROUP BY node_id",
		cs.Username)
	ir := PullData("lzy", influxsql)
	nodeIDList := ir.SortInfluxResp("desc")
	nodeIDList=nodeIDList[:cs.Step]
	err := DrainNodes(cs.Step, cs.ManagerIP, nodeIDList)
	if err != nil {
		glog.Error("ReduceCluster error drain node error", err)
		return err
	}
	//give sometime, containers need to rebuild
	time.Sleep(10 * time.Second)
	nodeIPs,err := cs.LeaveCluster(nodeIDList)
	if err != nil {
		glog.Error("Leave Cluster error", err)
		return err
	}
	//node info update need some time
	time.Sleep(15 * time.Second)
	err = DeleteNodes(cs.ManagerIP,nodeIDList)
	if err != nil {
		glog.Error("Delete Node error", err)
		return err
	}
	//if !cs.UpdateCluster(){
	//	return fmt.Errorf("updateCluster error")
	//}
	err=UpdateNodes(nodeIPs)
	return err
}

//choose a node from database
//add node to cluster
//update some service sort by cpu usage

//according to swarmid
func (cs ClusterStat) ListServiceC() []map[string]interface{} {
	sql := fmt.Sprintf("SELECT * FROM `service` WHERE swarm_id='%s'", cs.SwarmID)
	record := models.MysqlQuery(sql)
	return record
}

func (cs ClusterStat) CombineData(tempCpu int, tempMem int) compareData {
	return compareData{tempCpu, tempMem, cs.CpuUpper, cs.CpuLower,
		cs.MemUpper, cs.MemLower, cs.SwarmID}
}

//noList is ip array
func (cs ClusterStat) CheckExpand() ([]string, bool) {
	temp := clusterMap.Get(cs.SwarmID)
	if temp == nil {
		glog.Error("cluster node get error")
		return []string{}, false
	}
	nodeAmount, ok := temp.(int)
	NewNode, err := models.SelectNode(cs.Username, cs.Step)
	if err != nil {
		glog.Error("SelectNode error ",err)
		return NewNode, false
	}
	if len(NewNode) < cs.Step {
		glog.Error("Available nodes < Step")
		return NewNode, false
	}

	if ok && (nodeAmount+cs.Step <= cs.UpperLimit) {
		glog.Info("Expand check succeed")
		return NewNode, true
	}
	glog.Error("Expand check failed,node+step cannot > UpperLimit")
	return NewNode, false
}

//check cluster nodes
//check nodes after reduce
func (cs ClusterStat) CheckReduce() bool {
	temp := clusterMap.Get(cs.SwarmID)
	if temp == nil {
		glog.Error("cluster node get error")
		return false
	}
	nodeAmount, ok := temp.(int)
	if ok && (nodeAmount-cs.Step >= cs.LowerLimit) {
		glog.Info("Reduce check succeed")
		return true
	}
	glog.Error("Reduce check failed,node cannot < LowerLimit")
	return false
}

func (cs ClusterStat) LeaveCluster(nodeList []string) ([]string,error) {
	nodeIPs:=[]string{}
	for _, node := range nodeList {
		ip := GetNodeInfo(cs.ManagerIP, node).Status.Addr
		url := fmt.Sprintf("http://%s:2375/swarm/leave", ip)
		code, body, err := models.MyPost(url, []byte{})
		if err != nil || code > 200 {
			glog.Error("Leave Node error", err, string(body))
			return nodeIPs,err
		}
		nodeIPs=append(nodeIPs, ip)
	}
	glog.Info("Leave Node OK")
	return nodeIPs,nil
}

func (cs ClusterStat) JoinCluster(nodeList []string) bool {
	for _, ip := range nodeList {
		//add x node until x>step
		//if i >= cs.Step {
		//	glog.Info("join ", cs.Step, " nodes already")
		//	break
		//}

		cj := models.ClusterJoin{"0.0.0.0:2377", ip,
			cs.Token, []string{cs.ManagerIP + ":2377"}}

		sendCJ, err := json.Marshal(cj)
		if err != nil {
			glog.Error("sendCJ combine error", err)
			return false
		}
		url := fmt.Sprintf("http://%s:2375/swarm/join", ip)
		glog.Info(url)
		code, body, err := models.MyPost(url, sendCJ)
		if err != nil || code > 200 {
			glog.Error("swarm join failed", code, string(body))
			glog.Error(err)
			return false
		}
	}
	glog.Info("swarm join successed ")
	return true
}

func (cs ClusterStat) UpdateServices(serviceID []string) error{
	for _, id := range serviceID {
		glog.Info("Updating service ",id)
		si := models.GetServiceInfo(cs.ManagerIP, id)
		glog.V(1).Info(si.Spec)
		code, body, err := models.UpdateServiceDo(si.Spec, cs.ManagerIP, id, si.Version.Index)
		glog.Info(code, string(body), err)
		if err!=nil{
			return err
		}
	}
	return nil
}

func (cs ClusterStat)UpdateCluster() bool{
	url:=fmt.Sprintf("http://%s:2375/nodes",cs.ManagerIP)
	code, body, err := models.MyGet(url, map[string]string{})
	var nodeInfos []NodeInfo
	if err != nil || code > 200 {
		glog.Error("GetNode error")
		return false
	}
	json.Unmarshal(body, &nodeInfos)
	glog.V(1).Info(nodeInfos)
	for _,nodeInfo:=range nodeInfos{
		sql:=fmt.Sprintf("UPDATE `vm_info` SET `node_id`= '%s' , `hostname` = '%s'," +
			" `swarm_id` = '%s', `role` = '%s', `state` = '%s', `Availability` = '%s'" +
			" WHERE `vm_info`.`ip` = '%s';",
			nodeInfo.ID,nodeInfo.Description.Hostname,cs.SwarmID,nodeInfo.Role,nodeInfo.Status.State,
			nodeInfo.Availability,nodeInfo.Status.Addr)
		models.MysqlUpdate(sql)
	}
	return true

}