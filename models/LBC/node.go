package LBC

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
)

type NodeInfo struct {
	ID      string `json:"ID"`
	Version struct {
		Index int `json:"Index"`
	} `json:"Version"`
	NodeSpec `json:"Spec"`
	Description struct {
		Hostname string `json:"Hostname"`
	} `json:"Description"`
	Status   struct {
		State string `json:"State"`
		Addr  string `json:"Addr"`
	} `json:"Status"`
}
type NodeSpec struct {
	Name         string `json:"Name"`
	Role         string `json:"Role"`
	Availability string `json:"Availability"`
}

func DrainNodes(step int, managerIP string, nodeList []string) error {
	for _, node := range nodeList {
		//if i >= step {
		//	glog.Info("drain ", step, " nodes already")
		//	break
		//}
		version := GetNodeInfo(managerIP, node).Version.Index
		url := fmt.Sprintf("http://%s:2375/nodes/%s/update?version=%d", managerIP, node, version)
		var ns NodeSpec
		ns.Role = "worker"
		ns.Availability = "drain"

		sendNS, ok := json.Marshal(ns)
		if ok != nil {
			glog.Error("send_ns combine error")
		}
		code, body, err := models.MyPost(url, sendNS)
		if err != nil || code > 200 {
			glog.Error(url)
			glog.Error("Drain Node error", err, string(body))
			return err
		}
	}
	glog.Info("Drain Node OK")
	return nil

}

func GetNodeInfo(managerIP, nodeID string) NodeInfo {
	url := fmt.Sprintf("http://%s:2375/nodes/%s", managerIP, nodeID)
	code, body, err := models.MyGet(url, map[string]string{})
	var nodeInfo NodeInfo
	if err != nil || code > 200 {
		glog.Error("GetNodeVersion error")
	}
	json.Unmarshal(body, &nodeInfo)
	return nodeInfo
}

func DeleteNodes(managerIP string, nodeList []string) error{
	for _, node := range nodeList {
		url := fmt.Sprintf("http://%s:2375/nodes/%s", managerIP, node)
		code, body, err := models.MyDelete(url)
		if err != nil || code > 200 {
			glog.Error(url)
			glog.Error("Delete Node error", err, string(body))
			return err
		}
	}
	glog.Info("Delete Node OK")
	return nil
}

func UpdateNodes(nodeIPs []string)(error){
	var err error
	for _, nodeIP := range nodeIPs {
		sql:=fmt.Sprintf("UPDATE `vm_info` SET `node_id`= NULL ," +
			" `swarm_id` = NULL , `role` = NULL, `state` = NULL, `Availability` = NULL" +
			" WHERE `vm_info`.`ip` = '%s';", nodeIP)
		_,_,err=models.MysqlUpdate(sql)
		if err!=nil {
			return err
		}
	}
	return err
}
