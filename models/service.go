package models

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
)

type ServiceInfo struct {
	ID      string `json:"ID"`
	Version struct {
		Index int `json:"Index,omitempty"`
	} `json:"Version,omitempty"`
	Spec ServiceSpec `json:"Spec"`
}

type ServiceSpec struct {
	Name         string `json:"Name"`
	TaskTemplate struct {
		ContainerSpec struct {
			Image string `json:"Image"`
		} `json:"ContainerSpec"`
		Placement struct {
			Constraints []string `json:"Constraints"`
		} `json:"Placement"`
		Networks []struct {
			Target string `json:"Target"`
		} `json:"Networks"`
		ForceUpdate int `json:"ForceUpdate"`
	} `json:"TaskTemplate"`
	Mode struct {
		Replicated struct {
			Replicas int `json:"Replicas"`
		} `json:"Replicated"`
	} `json:"Mode"`
	Labels struct {
		TraefikPort         string `json:"traefik.port"`
		TraefikFrontendRule string `json:"traefik.frontend.rule"`
	} `json:"Labels"`
}
type FrontendSS struct {
	Name                string `json:"Name"`
	Image               string `json:"Image"`
	SwarmID             string `json:"swarm_id"`
	Constraints         string `json:"Constraints"`
	Target              string `json:"Target"`
	Replicas            int    `json:"Replicas"`
	TraefikPort         string `json:"traefik.port"`
	TraefikFrontendRule string `json:"traefik.frontend.rule"`
	ElasticInfo
}

type MyResponse struct {
	Code       int    `json:"res_code"`
	Body       []byte `json:"res_body"`
	Err        error  `json:"err"`
	LastInsert int64  `json:"last_insert"`
	RowsAffect int64  `json:"rows_affect"`
}

//according to uid
func ListServiceU(uid int) []map[string]interface{} {
	sql := fmt.Sprintf("SELECT * FROM `service` WHERE uid=%d", uid)
	record := MysqlQuery(sql)
	return record
}

func GetServiceInfo(managerIP, serviceID string) ServiceInfo {
	url := fmt.Sprintf("http://%s:2375/services/%s", managerIP, serviceID)
	code, body, err := MyGet(url, map[string]string{})
	var serviceInfo ServiceInfo
	if err != nil || code > 200 {
		glog.Error("GetServiceInfo error")
	}
	glog.V(2).Info(string(body))
	json.Unmarshal(body, &serviceInfo)
	return serviceInfo
}

func UpdateServiceDo(data ServiceSpec, managerIP string, serviceID string, version int) (int, []byte, error) {
	//to force update Service we must add this value,
	// so it is different from previous one
	data.TaskTemplate.ForceUpdate++
	sendData, err := json.Marshal(data)
	glog.Info(string(sendData))
	if err != nil {
		glog.Error("sendData combine error", err)
		return 400, []byte{}, err
	}
	url := fmt.Sprintf("http://%s:2375/services/%s/update?version=%d", managerIP, serviceID, version)

	glog.Info(url)
	code, body, err := MyPost(url, sendData)
	if err != nil || code > 200 {
		glog.Error("Update service failed ", code, string(body))
		glog.Error(err)
	}
	return code, body, err
}
