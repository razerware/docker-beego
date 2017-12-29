package controllers

import "github.com/astaxie/beego"

type ServiceController struct {
	beego.Controller
}

func (c *ServiceController) ServiceApply() {
	c.TplName = "service_apply.tpl"
}

func (c *ServiceController) ServiceDetail() {
	c.TplName = "service_detail.tpl"
}

func (c *ServiceController) ServiceMonitor() {
	c.TplName = "service_monitor.tpl"
}
