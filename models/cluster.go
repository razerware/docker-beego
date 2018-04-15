package models

import (
	"fmt"
	"github.com/golang/glog"
)

type ClusterInfo []struct {
	Swarm_ID   int    `json:"swarm_id"`
	Token      string `json:"token"`
	Manager_ip string `json:"manager_ip"`
	User
	ElasticInfo
}

type ClusterInit struct {
	//ListenAddr      string `json:"ListenAddr"`
	AdvertiseAddr string `json:"AdvertiseAddr"`
	Spec struct {
		Name string `json:"Name"`
	} `json:"Spec"`
}

type ClusterJoin struct {
	ListenAddr    string   `json:"listenAddr"`
	AdvertiseAddr string   `json:"AdvertiseAddr"`
	JoinToken     string   `json:"JoinToken"`
	RemoteAddrs   []string `json:"RemoteAddrs"`
}

type FrontendCI struct {
	SwarmId       string `json:"SwarmId"`
	AdvertiseAddr string `json:"AdvertiseAddr"`
	Name          string `json:"Name"`
	ElasticInfo
}

type FrontendCJ struct {
	//ListenAddr    string `json:"listenAddr"`
	AdvertiseAddr string `json:"AdvertiseAddr"`
	JoinToken     string `json:"JoinToken"`
	ManagerIp     string `json:"ManagerIp"`
}

type FrontendJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func RegistCluster(fci FrontendCI, uid int) {

	//c.Ctx.Input.Bind(&uid, "uid")
	//json.Unmarshal(c.Ctx.Input.RequestBody, &fci)
	sql := fmt.Sprintf("INSERT INTO `cluster_info` (`swarm_id`,`name`,`uid`,`manager_ip`,"+
		"`upper_limit`, `lower_limit`, `step`,`cpu_lower`, `cpu_upper`, `mem_lower`, `mem_upper`) "+
		"VALUES ('%s','%s','%d','%s','%d', '%d', '%d', '%d', '%d', '%d', '%d')",
		fci.SwarmId, fci.Name, uid, fci.AdvertiseAddr,
		fci.UpperLimit, fci.LowerLimit, fci.Step, fci.CpuLower, fci.CpuUpper, fci.MemLower, fci.MemUpper)
	fmt.Println(sql)
	last, row, err := MysqlInsert(sql)
	glog.Info(last, row, err)

}

func ListCluster(uid int) []map[string]interface{}{
	sql := fmt.Sprintf("SELECT * FROM `cluster_info` WHERE uid=%d", uid)
	record := MysqlQuery(sql)
	return record
}