package controllers

import (
	"fmt"
	"github.com/razerware/docker_beego/models"
	"encoding/json"
	"github.com/golang/glog"
)

type ClusterController struct {
	BaseController
}

// @router /cluster_detail [get]
func (c *ClusterController) ClusterDetail() {
	c.TplName = "cluster_detail.tpl"
}

// @router /cluster_apply [get]
func (c *ClusterController) ClusterApply() {
	c.TplName = "cluster_apply.tpl"
}

// @router /cluster_apply [post]
func (c *ClusterController) ClusterInit() {
	var manager_ip, name string
	c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	c.Ctx.Input.Bind(&name, "name")
	var fci models.FrontendCI
	json.Unmarshal(c.Ctx.Input.RequestBody, &fci)
	glog.Info(fci)
	var ci models.ClusterInit
	ci.AdvertiseAddr = manager_ip
	ci.Spec.Name = name

	send_ss, ok := json.Marshal(ci)
	if ok != nil {
		glog.Error(ok)
	}
	url := "http://" + manager_ip + ":2375/swarm/init"
	code, body, err := models.MyPost(url, send_ss)
	if err != nil || code > 200 {
		glog.Error(code)
		glog.Error(err)
	}
	glog.Info(string(send_ss))
	glog.Info(body)
	fci.SwarmId = string(body)
	models.RegistCluster(fci, c.userId)

	c.Data["json"] = ci
	c.ServeJSON()
}

// @router /cluster_join [post]
func (c *ClusterController) ClusterJoin() {

	var fcj models.FrontendCJ
	json.Unmarshal(c.Ctx.Input.RequestBody, &fcj)
	var cj models.ClusterJoin
	cj.ListenAddr = "0.0.0.0:2377"
	cj.AdvertiseAddr = fcj.AdvertiseAddr
	cj.JoinToken = fcj.JoinToken
	cj.RemoteAddrs = append(cj.RemoteAddrs, fcj.ManagerIp+":2377")
	sendCJ, ok := json.Marshal(cj)
	if ok != nil {
		glog.Error(ok)
		c.Data["json"] = ""
		c.ServeJSON()
	}
	glog.Info(string(sendCJ), ok)
	url := "http://" + fcj.AdvertiseAddr + ":2375/swarm/join"
	glog.Info(url)
	code, body, err := models.MyPost(url, sendCJ)
	if err != nil || code > 200 {
		glog.Error("swarm join failed", code)
		glog.Error(err)
		c.Data["json"] = fcj
		c.ServeJSON()
	}
	glog.Info("swarm join successed ", string(sendCJ))
	glog.Info("resp is ", string(body))
	c.Data["json"] = fcj
	c.ServeJSON()
}

// @router /cluster_list [get]
func (c *ClusterController) ClusterList() {
	sql := fmt.Sprintf("SELECT * FROM cluster_info WHERE uid=%d", c.userId)
	record := models.MysqlQuery(sql)
	fj := models.FrontendJson{0, "", len(record), record}
	c.Data["json"] = fj
	c.ServeJSON()
}

// @router /cluster_expand [post]
func (c *ClusterController) ClusterExpand(managerIp string, ip string, joinToken string, step int) (bool, error) {
	//var manager_ip,name string
	//c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	sql := "SELECT ip FROM `vm_info` WHERE `uid` = 1 AND `swarm_id` = ''"
	record := models.MysqlQuery(sql)
	if len(record) <= step {
		for _, v := range record {
			c.ClusterJoin()
			fmt.Println(v)
		}
	}

	return false, nil
}
