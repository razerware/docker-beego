package controllers

import (
    "github.com/astaxie/beego"
)

type HdrController struct {
    beego.Controller
}

func (c *HdrController) Get() {
    c.TplName = "hdr.tpl"
}
