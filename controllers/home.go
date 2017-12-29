package controllers

import (
    "github.com/astaxie/beego"
)

type HomeController struct {
    beego.Controller
}

func (c *HomeController) Get() {
    // 从Router拿到username
    c.username = this.Ctx.Input.Param(":username")

    c.Ctx.SetCookie("username", c.username)
    if(c.username === 'lzy'){
        c.Ctx.SetCookie("url", '10.109.252.172:9999')
    } else {
        c.Ctx.SetCookie("url", '10.109.252.172:10000')
    }
    c.TplName = "layout.tpl"
}





