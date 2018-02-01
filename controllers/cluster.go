package controllers

import (
	"github.com/astaxie/beego"
	"docker-beego/elbController"
)

type ClusterController struct {
	beego.Controller
}

type ClusterInfo struct {
	ID int
	Token string `json:"token"`
	User
	elbController.ElasticInfo
	manager_ip string
}
func (c *ClusterController) ClusterMonitor() {
	c.TplName = "cluster_monitor.tpl"
}

func (c *ClusterController) ClusterDetail() {
	c.TplName = "cluster_detail.tpl"
}

func (c *ClusterController) ClusterApply() {
	c.TplName = "cluster_apply.tpl"
}

func (c *ClusterController) Test() {
	c.TplName = "cluster_apply.tpl"
}