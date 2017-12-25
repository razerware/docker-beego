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

func (c *HdrController) ClusterInfo() {
    c.TplName = "cluster_info.tpl"
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

func (c *HdrController) ServiceApply() {
    c.TplName = "service_apply.tpl"
}

func (c *HdrController) ServiceDetail() {
    c.TplName = "service_detail.tpl"
}

func (c *HdrController) ServiceInfo() {
    c.TplName = "service_info.html"
}
