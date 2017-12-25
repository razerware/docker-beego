package controllers

import (
    "github.com/astaxie/beego"
)

type HdrController struct {
    beego.Controller
}

func (c *HdrController) Get() {
    c.TplName = "all.tpl"
}

func (c *HdrController) Test() {
    c.TplName = "test.tpl"
}

func (c *HdrController) Test2() {
    c.TplName = "test2.tpl"
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
