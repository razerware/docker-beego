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
