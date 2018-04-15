package controllers

import (
	"fmt"
	"github.com/razerware/docker_beego/models"
	"encoding/json"
	"github.com/golang/glog"
)

type ServiceController struct {
	BaseController
}

// @router /service_monitor [get]
func (c *ServiceController) ServiceMonitor() {
	c.TplName = "service_monitor.tpl"
}

func (c *ServiceController) ServiceUpdate(id string) {
	c.TplName = "service_update.tpl"
}

func (c *ServiceController) ServiceUpdatePost(id string) {
}

// @router /service_apply [get]
func (c *ServiceController) ServiceApply() {
	c.TplName = "service_apply.tpl"
}

// @router /service_apply [post]
func (c *ServiceController) ServiceApplyPost() {
	uid := c.userId
	//uid=1
	//get manager_ip of a swarm from frontend
	var manager_ip string
	c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	var fss models.FrontendSS
	json.Unmarshal(c.Ctx.Input.RequestBody, &fss)
	glog.Info(string(c.Ctx.Input.RequestBody))
	glog.Info(fss)
	//transfer frontend json to backend readable struct ServiceSpec
	var ss models.ServiceSpec
	ss.Name = fss.Name
	ss.TaskTemplate.ContainerSpec.Image = fss.Image
	ss.TaskTemplate.Placement.Constraints = append(ss.TaskTemplate.Placement.Constraints, fss.Constraints)
	ss.TaskTemplate.Networks = append(ss.TaskTemplate.Networks, struct{ Target string `json:"Target"` }{fss.Target})
	ss.Mode.Replicated.Replicas = fss.Replicas
	ss.Labels.TraefikPort = fss.TraefikPort
	//ss.Labels.TraefikFrontendRule = fss.TraefikFrontendRule
	ss.Labels.TraefikFrontendRule = fss.Name + ".com"
	//transfer ServiceSpec to json for Docker API
	send_ss, ok := json.Marshal(ss)
	if ok != nil {
		glog.Error(ok)
	}
	//combine url and post request to DockerAPI
	url := "http://" + manager_ip + ":2375/services/create"
	var res models.MyResponse
	res.Code, res.Body, res.Err = models.MyPost(url, send_ss)
	if res.Err != nil || res.Code > 201 {
		//Error occurs so stop here
		glog.Error(res.Code)
		glog.Error(string(res.Body))
		c.Data["json"] = res
		c.CustomAbort(res.Code, string(res.Body))
		//c.ServeJSON()
		//c.StopRun()
	}
	glog.V(1).Info(send_ss)
	glog.V(1).Info(string(res.Body))

	//Get response from DockerAPI,unmarhal ID(service_id) to a map
	m := make(map[string]string)
	res.Err = json.Unmarshal([]byte(res.Body), &m)
	if serviceId, ok := m["ID"]; ok {
		//id exist so insert info to MysqlDB
		glog.Info("Service " + serviceId + " create achived")
		sql := fmt.Sprintf("INSERT INTO `service` (`service_id`,`name`,`uid`,"+
			"`desire_replica`,`swarm_id`,`image`"+
			"`upper_limit`, `lower_limit`, `step`,`cpu_lower`, `cpu_upper`, `mem_lower`, `mem_upper`) "+
			"VALUES ('%s','%s','%d','%d','%s','%d','%s', '%d', '%d', '%d', '%d', '%d', '%d')",
			serviceId, fss.Name, uid, fss.Replicas, fss.SwarmID, fss.Image, fss.UpperLimit, fss.LowerLimit, fss.Step,
			fss.CpuLower, fss.CpuUpper, fss.MemLower, fss.MemUpper)
		glog.V(1).Info(sql)
		res.LastInsert, res.RowsAffect, res.Err = models.MysqlInsert(sql)
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// @router /service_detail [get]
func (c *ServiceController) ServiceDetail() {
	c.TplName = "service_detail.tpl"
}

// @router /list_service [get]
func (c *ServiceController) ListService() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	if v, ok := uid.(int); ok {
		uid = v
	} else {
		uid = 0
	}
	//var uid int
	//c.Ctx.Input.Bind(&uid, "uid")
	sql := fmt.Sprintf("SELECT * FROM `service` WHERE uid=%d", uid)
	record := models.MysqlQuery(sql)
	fj := models.FrontendJson{0, "", len(record), record}
	c.Data["json"] = fj
	c.ServeJSON()

}
