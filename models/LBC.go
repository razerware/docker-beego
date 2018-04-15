package models

import (
	"github.com/golang/glog"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego"
	"fmt"
	"strings"
)

var (
	//MontiorMap map[string]Monitor
	monitorMap = utils.NewBeeMap()
	clusterMap = utils.NewBeeMap()
)

type ElasticInfo struct {
	UpperLimit int `json:"upper_limit"`
	LowerLimit int `json:"lower_limit"`
	Step       int `json:"step"`
	CpuLower   int `json:"cpu_lower"`
	CpuUpper   int `json:"cpu_upper"`
	MemLower   int `json:"mem_lower"`
	MemUpper   int `json:"mem_upper"`
}

type ClusterStat struct {
	SwarmId string `json:"swarm_id"`
	ElasticInfo
}

type ServiceStat struct {
	ServiceId     string `json:"service_id"`
	SwarmId       string `json:"swarm_id"`
	DesireReplica int    `json:"desire_replica"`
	ElasticInfo
}

type Monitor struct {
	CpuFlag int
	MemFlag int
}

type InfluxResp struct {
	Results []struct {
		StatementID int `json:"statement_id"`
		Series []struct {
			Name    string          `json:"name"`
			Columns []string        `json:"columns"`
			Values  [][]interface{} `json:"values"`
		} `json:"series"`
	} `json:"results"`
}

type compareData struct {
	tempCpu  int
	tempMem  int
	cpuUpper int
	cpuLower int
	memUpper int
	memLower int
	mapKey   string
}

func DataCollecter() {
	sleep := make(chan int)
	go func() {
		time.Sleep(30 * time.Second)
		sleep <- 1
	}()
	for {
		select {
		case <-sleep:
			userList := ListUser()
			CountClusterNode()
			for _, user := range userList {
				CheckService(user)
				//go Judgement(data, user.Username)
			}
		}
	}
}

func CountClusterNode() {
	sql := "SELECT COUNT(*)AS amount,swarm_id FROM `vm_info` WHERE swarm_id!='' GROUP BY swarm_id"
	record := MysqlQuery(sql)
	if record != nil {
		for _, i := range record {
			glog.Info(i["swarm_id"])
			clusterMap.Set(i["swarm_id"], i["amount"])
		}
		glog.Info(clusterMap)
	}
}

func CheckService(user User) {
	record := ListService(user.Uid)
	if record == nil {
		glog.Error("No services")
		return
	}
	var ss []ServiceStat
	//map->json->struct
	if temp, ok := json.Marshal(record); ok == nil {
		glog.Info(string(temp))
		ok = json.Unmarshal(temp, &ss)
		glog.Info(ok)
		glog.Info(ss)
	} else {
		glog.Error(ok)
	}
	for _, s := range ss {
		CpuSql := InfluxSelect("sum", "cpu", user.Username, "autogen", "container",
			"30s", "serviceid", s.ServiceId)
		MemSql := InfluxSelect("mean", "mem", user.Username, "autogen", "vm",
			"30s", "serviceid", s.ServiceId)
		ir := PullData(user.Username, CpuSql+";"+MemSql)
		result := Judgement(ir, s)
		amount := clusterMap.Get(s.SwarmId)
		glog.Info(result, amount)
		//if s.DesireReplica<nodeNumber{
		//
		//}
	}

}

func CheckCluster(user User) {
	record := ListCluster(user.Uid)
	if record == nil {
		glog.Error("No Clusters")
		return
	}
	var cs []ClusterStat
	//map->json->struct
	if temp, ok := json.Marshal(record); ok == nil {
		glog.Info(string(temp))
		ok = json.Unmarshal(temp, &cs)
		glog.Info(ok)
		glog.Info(cs)
	} else {
		glog.Error(ok)
	}
	for _, s := range cs {
		CpuSql := InfluxSelect("mean", "cpu", user.Username, "autogen", "vm",
			"30s", "swarmid", s.SwarmId)
		MemSql := InfluxSelect("mean", "mem", user.Username, "autogen", "vm",
			"30s", "swarmid", s.SwarmId)
		ir := PullData(user.Username, CpuSql+";"+MemSql)
		result := Judgement(ir, s)
		amount := clusterMap.Get(s.SwarmId)
		glog.Info(result, amount)
		//if s.DesireReplica<nodeNumber{
		//
		//}
	}

}

func InfluxSelect(function, field, db, policy, measurement, period, specK, specV string) string {
	sql := fmt.Sprintf("SELECT %s(%s) AS %s_%s FROM %s.%s.%s WHERE time > now() - %s AND %s='%s'",
		function, field, function, field, db, policy, measurement, period, specK, specV)
	return sql
}

func PullData(user string, sql string) InfluxResp {
	paras := map[string]string{}
	var ir InfluxResp
	paras["pretty"] = "true"
	paras["db"] = user
	paras["q"] = sql
	code, body, err := MyGet("http://"+beego.AppConfig.String("InfluxUrl")+"/query", paras)
	//code, body, err := MyGet(beego.AppConfig.String("InfluxUrl"), paras)
	if err != nil || code > 200 {
		glog.Error(err)
		return ir
	}
	glog.Info(string(body))
	json.Unmarshal(body, &ir)
	return ir
}

func Judgement_abandon(ir InfluxResp, ss ServiceStat) bool {
	if ir.Results[0].Series == nil {
		glog.Error("Cpu data error")
		return false
	}
	v, _ := ir.Results[0].Series[0].Values[0][1].(float64)
	tempCpu := int(v)
	if ir.Results[1].Series == nil {
		glog.Error("Mem data error")
		return false
	}
	v, _ = ir.Results[1].Series[0].Values[0][1].(float64)
	tempMem := int(v)
	glog.Info(tempCpu, tempMem)

	vMontior := monitorMap.Get(ss.ServiceId)
	monitor, _ := vMontior.(Monitor)
	monitor.CpuFlag = AlgoCompare(tempCpu, ss.CpuUpper, ss.CpuLower, monitor.CpuFlag)
	monitor.MemFlag = AlgoCompare(tempMem, ss.MemUpper, ss.MemLower, monitor.MemFlag)
	DynamicPeriod, _ := beego.AppConfig.Int("DynamicPeriod")
	periodFlag := 2<<uint(DynamicPeriod) - 1

	if monitor.CpuFlag >= periodFlag || monitor.MemFlag >= periodFlag {
		//ss.DesireReplica
		if monitor.CpuFlag > 0 {
			//give it some cd
			monitor.CpuFlag = 0
		} else {
			monitor.MemFlag = 0
		}
		monitorMap.Set(ss.ServiceId, monitor)
		return true
	}
	return false
}

func Judgement(ir InfluxResp, statInterface interface{}) int {
	if len(ir.Results) < 1 {
		glog.Error("InfluxResp not contain both CPU or MEM data")
		return 0
	}
	var tempCpu, tempMem int
	for _, i := range ir.Results {
		if len(i.Series) == 0 {
			glog.Error("Need both Cpu and Mem data from InfluxDB")
			continue
		}
		if strings.Contains(i.Series[0].Columns[1], "cpu") {
			v, _ := i.Series[0].Values[0][1].(float64)
			tempCpu = int(v)
			glog.V(1).Info("this is Cpu data:", tempCpu)
		} else if strings.Contains(i.Series[0].Columns[1], "mem") {
			v, _ := i.Series[0].Values[0][1].(float64)
			tempMem = int(v)
			glog.V(1).Info("this is Mem data", tempMem)
		}
	}
	var vMonitor interface{}
	switch stat := statInterface.(type) {
	case ServiceStat:
		glog.V(1).Info("check Service")
		vMonitor = monitorMap.Get(stat.ServiceId)
		cd := compareData{tempCpu, tempMem, stat.CpuUpper, stat.CpuLower,
			stat.MemUpper, stat.MemLower, stat.ServiceId}
		return JudgeByData(vMonitor, cd)
	case ClusterStat:
		glog.V(1).Info("check Cluster")
		vMonitor = monitorMap.Get(stat.SwarmId)
		cd := compareData{tempCpu, tempMem, stat.CpuUpper, stat.CpuLower,
			stat.MemUpper, stat.MemLower, stat.SwarmId}
		return JudgeByData(vMonitor, cd)
	default:
		return 0
	}
}

func JudgeByData(vMonitor interface{}, data compareData) int {
	monitor, _ := vMonitor.(Monitor)
	monitor.CpuFlag = AlgoCompare(data.tempCpu, data.cpuUpper, data.cpuLower, monitor.CpuFlag)
	monitor.MemFlag = AlgoCompare(data.tempMem, data.memUpper, data.memLower, monitor.MemFlag)
	DynamicPeriod, _ := beego.AppConfig.Int("DynamicPeriod")
	periodFlag := 1<<uint(DynamicPeriod) - 1
	glog.V(1).Info("After compare monitor is: ", monitor)
	//expand is "or" but reduce need "and"
	result:=0
	if monitor.CpuFlag >= periodFlag || monitor.MemFlag >= periodFlag {
		//expand/reduce need some CD so they cannot operated together
		monitor.CpuFlag = 0
		monitor.MemFlag = 0
		glog.V(1).Info("Need Expand")
		result=1
	}else if monitor.CpuFlag <= -periodFlag && monitor.MemFlag <= -periodFlag {
		monitor.CpuFlag = 0
		monitor.MemFlag = 0
		glog.V(1).Info("Need reduce")
		result=-1
	}else {
		glog.V(1).Info("No Need to expand/reduce")
		result=0
	}
	//flag cannot exceed periodFlag
	monitor.CpuFlag=thresholdFlag(monitor.CpuFlag,periodFlag)
	monitor.MemFlag=thresholdFlag(monitor.MemFlag,periodFlag)
	monitorMap.Set(data.mapKey, monitor)
	glog.V(1).Info(monitorMap.Items())
	glog.V(1).Info("Judgement complate False")
	return result
}

func AlgoCompare(test int, upper int, lower int, flag int) int {

	if test > upper {
		if flag > 0 {
			flag = flag<<1 + 1
		} else {
			flag = 1
		}
	} else if test < lower {
		if flag < 0 {
			flag = flag<<1 - 1
		} else {
			flag = -1
		}
	} else {
		flag = 0
	}
	return flag
}

func JudgeVm(data chan int, threshold int, resp chan bool) bool {
	over := 0
	t := 0
	count := 0
	list := make([]int, 10)
	for {
		select {
		case i := <-data:
			if i > threshold {
				over = 1
			} else {
				over = 0
			}
			count = (count + 1) % 11
			t = t - list[count] + over
			list[count] = over
			if t > 7 {
				//ClusterJoin()
				resp <- true
				glog.Info("need Cluster Expand")
				t = 0
			}
		}
	}
}

func JudgeContainer(data chan int, threshold int, resp chan bool) bool {
	t := 0
	for {
		select {
		case i := <-data:
			if i > threshold {
				t++
			} else {
				t = 0
			}
			if t == 5 {
				//ClusterJoin()
				resp <- true
				glog.Info("need Cluster Expand")
				t = 0
			}
		}
	}
}

func Lzytest(a interface{}) Monitor {
	v, _ := a.(Monitor)
	glog.Info(v)
	return v
}

func myAbs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func thresholdFlag(input int,threshold int) int{
	if input>threshold{
		return threshold
	}
	if input< -threshold{
		return -threshold
	}
	return input
}