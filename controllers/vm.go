package controllers

import "github.com/astaxie/beego"

type VmController struct {
	beego.Controller
}
func (c *VmController) VmApply() {
	c.TplName = "vm_apply.tpl"
}

func (c *VmController) VmDetail() {
	c.TplName = "vm_detail.tpl"
}

func (c *VmController) VmDetailApi() {
	c.TplName = "vm_detail_api.tpl"
}