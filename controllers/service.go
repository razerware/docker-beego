package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type ServiceController struct {
	beego.Controller
}

type LBspec []struct {
	Address     string
	Swarm_ID     int
	Image       string
	Upper_limit int
	Lower_limit int
	Step        int
	Cpu_lower   int
	Cpu_upper   int
	Mem_lower   int
	Mem_upper   int
}

type ServiceStat []struct {
	ID string `json:"ID"`
	Spec struct {
		Name string `json:"Name"`
		Labels struct {
			TraefikFrontendRule string `json:"traefik.frontend.rule"`
			TraefikPort         string `json:"traefik.port"`
		} `json:"Labels"`
		TaskTemplate struct {
			ContainerSpec struct {
				Image string `json:"Image"`
			} `json:"ContainerSpec"`
		} `json:"TaskTemplate"`
		Mode struct {
			Replicated struct {
				Replicas int `json:"Replicas"`
			} `json:"Replicated"`
		} `json:"Mode"`
	} `json:"Spec"`
}

func (c *ServiceController) ServiceApply() {
	c.TplName = "service_apply.tpl"
}

func (c *ServiceController) ServiceDetail(cluster_info ClusterInfo) {
	c.TplName = "service_detail.tpl"

	client := &http.Client{}
	url := "http://" + cluster_info.manager_ip + ":2375/services"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	ss := []ServiceStat{}
	json.Unmarshal(body, &ss)
	fmt.Println(ss)
}

func (c *ServiceController) ListService() {
	var manager_ip string
	c.Ctx.Input.Bind(&manager_ip, "manager_ip")

	client := &http.Client{}
	url := "http://" + manager_ip + ":2375/services"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	ss := ServiceStat{}
	json.Unmarshal(body, &ss)
	svc_name := []string{}
	for _, i := range ss {
		svc_name = append(svc_name, i.Spec.Name)
	}
	fmt.Println(svc_name)
	c.Data["json"] = svc_name
	c.ServeJSON()

}

func (c *ServiceController) ServiceMonitor() {
	c.TplName = "service_monitor.tpl"
}
