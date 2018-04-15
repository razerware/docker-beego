package models

import "fmt"

type ServiceInfo []struct {
	Service_ID string `json:"service_id"`
	Swarm_ID   int    `json:"swarm_id"`
	Address    string `json:"Address"`
	Replication int `json:"replication"`
	ServiceSpec
	ElasticInfo
}

type ServiceSpec struct {
	Name string `json:"Name"`
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
	Code       int `json:"res_code"`
	Body       []byte `json:"res_body"`
	Err        error `json:"err"`
	LastInsert int64 `json:"last_insert"`
	RowsAffect int64 `json:"rows_affect"`
}

func ListService(uid int) []map[string]interface{}{
	sql := fmt.Sprintf("SELECT * FROM `service` WHERE uid=%d", uid)
	record := MysqlQuery(sql)
	return record
}