package controllers

import (
    "github.com/astaxie/beego"
)

type HdrController struct {
    beego.Controller
}

func (c *HdrController) Get() {
    c.TplName = "layout.tpl"
}

func (c *HdrController) ClusterMonitor() {
    c.TplName = "cluster_monitor.tpl"
}

func (c *HdrController) ClusterDetail() {
    c.TplName = "cluster_detail.tpl"
}

func (c *HdrController) ClusterApply() {
    c.TplName = "cluster_apply.tpl"
}

func (c *HdrController) VmApply() {
    c.TplName = "vm_apply.tpl"
}

func (c *HdrController) VmDetail() {
    c.TplName = "vm_detail.tpl"
}

func (c *HdrController) VmDetailApi() {
    c.TplName = "vm_detail_api.tpl"
}

func (c *HdrController) ServiceApply() {
    c.TplName = "service_apply.tpl"
}

func (c *HdrController) ServiceDetail() {
    c.TplName = "service_detail.tpl"
}

func (c *HdrController) ServiceMonitor() {
    c.TplName = "service_monitor.tpl"
}
