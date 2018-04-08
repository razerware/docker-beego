package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"docker-beego/models"
	"docker-beego/elbController"
	"encoding/json"
	"github.com/golang/glog"
)

type ServiceController struct {
	beego.Controller
}

type ServiceInfo []struct {
	Service_ID string `json:"service_id"`
	Swarm_ID   int    `json:"swarm_id"`
	Address    string `json:"Address"`
	ServiceSpec
	elbController.ElasticInfo
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
	SwarmID             string `json:"SwarmId"`
	Constraints         string `json:"Constraints"`
	Target              string `json:"Target"`
	Replicas            int    `json:"Replicas"`
	TraefikPort         string `json:"traefik.port"`
	TraefikFrontendRule string `json:"traefik.frontend.rule"`
	elbController.ElasticInfo
}

type MyResponse struct {
	Code       int `json:"res_code"`
	Body       []byte `json:"res_body"`
	Err        error `json:"err"`
	LastInsert int64 `json:"last_insert"`
	RowsAffect int64 `json:"rows_affect"`
}

// @router /service_apply [get]
func (c *ServiceController) ServiceApply() {
	c.TplName = "service_apply.tpl"
}

// @router /service_apply [post]
func (c *ServiceController) ServiceApplyPost() {
	//username := c.Ctx.GetCookie("username")
	//uid := c.GetSession(username)
	uid:=1
	//get manager_ip of a swarm from frontend
	var manager_ip string
	c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	var fss FrontendSS
	json.Unmarshal(c.Ctx.Input.RequestBody, &fss)
	glog.Info(string(c.Ctx.Input.RequestBody))
	glog.Info(fss)
	//transfer frontend json to backend readable struct ServiceSpec
	var ss ServiceSpec
	ss.Name = fss.Name
	ss.TaskTemplate.ContainerSpec.Image = fss.Image
	ss.TaskTemplate.Placement.Constraints = append(ss.TaskTemplate.Placement.Constraints, fss.Constraints)
	ss.TaskTemplate.Networks = append(ss.TaskTemplate.Networks, struct{ Target string `json:"Target"` }{fss.Target})
	ss.Mode.Replicated.Replicas = fss.Replicas
	ss.Labels.TraefikPort = fss.TraefikPort
	ss.Labels.TraefikFrontendRule = fss.TraefikFrontendRule
	//transfer ServiceSpec to json for Docker API
	send_ss, ok := json.Marshal(ss)
	if ok != nil {
		glog.Error(ok)
	}
	//combine url and post request to DockerAPI
	url := "http://" + manager_ip + ":2375/services/create"
	var res MyResponse
	res.Code, res.Body, res.Err = MyPost(url, send_ss)
	if res.Err != nil || res.Code > 201 {
		//Error occurs so stop here
		glog.Error(res.Code)
		glog.Error(res.Err)
		c.Data["json"] = res
		c.ServeJSON()
		c.StopRun()
	}
	glog.V(1).Info(send_ss)
	glog.V(1).Info(string(res.Body))

	//Get response from DockerAPI,unmarhal ID(service_id) to a map
	m := make(map[string]string)
	res.Err = json.Unmarshal([]byte(res.Body), &m)
	if serviceId, ok := m["ID"]; ok {
		//id exist so insert info to MysqlDB
		glog.Info("Service "+serviceId+" create achived")
		sql := fmt.Sprintf("INSERT INTO `service` (`service_id`,`name`,`uid`,`swarm_id`,`upper_limit`, `lower_limit`, `step`, "+
			"`cpu_lower`, `cpu_upper`, `mem_lower`, `mem_upper`) "+
			"VALUES ('%s','%s','%d','%s','%d', '%d', '%d', '%d', '%d', '%d', '%d')",
			serviceId,fss.Name,uid,fss.SwarmID,fss.UpperLimit, fss.LowerLimit, fss.Step, fss.CpuLower, fss.CpuUpper, fss.MemLower, fss.MemUpper)
		glog.V(1).Info(sql)
		res.LastInsert, res.RowsAffect, res.Err = models.MysqlInsert(sql)
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ServiceController) ServiceDetail() {
	c.TplName = "service_detail.tpl"
}

func (c *ServiceController) ListService() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	//var uid int
	//c.Ctx.Input.Bind(&uid, "uid")
	sql := fmt.Sprintf("SELECT * FROM `service` WHERE uid=%d", uid)
	record := models.MysqlQuery(sql)
	fj:=FrontendJson{0,"",len(record),record}
	c.Data["json"] = fj
	c.ServeJSON()

}

func (c *ServiceController) ServiceMonitor() {
	c.TplName = "service_monitor.tpl"
}

func (c *ServiceController) ServiceUpdate(id string) {
	c.TplName = "service_update.tpl"
}

func (c *ServiceController) ServiceUpdatePost(id string) {
}
