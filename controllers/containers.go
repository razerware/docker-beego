package controllers

import (
"github.com/astaxie/beego"
)

type ContainerController struct {
	beego.Controller
}

func (c *ContainerController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *ContainerController) GetContainers() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}