package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
    "github.com/golang/glog"
)

type HomeController struct {
    beego.Controller
}

func (c *HomeController) Get() {
    // 从Router拿到username
    username:=c.Ctx.GetCookie("username")
    glog.V(1).Info(username)
    uid := c.GetSession(username)
    glog.V(1).Info(uid)
    if uid==nil{
        glog.Info("session not exist")
        c.Redirect(beego.URLFor("LoginController.Get"),302)
        c.StopRun()
    }
    if(username == "lzy"){
        c.Ctx.SetCookie("url", "10.109.252.172:8888")
    } else {
        c.Ctx.SetCookie("url", "10.109.252.172:10000")
    }
    c.TplName = "layout.tpl"
}

func (c *HomeController) GetUid() {
    // 从Router拿到username
    session_id:=c.Ctx.GetCookie("session_id")
    fmt.Println(session_id)
    v := c.GetSession(session_id)
    fmt.Println(v)
    if v==nil{
        fmt.Println("session not exist")
        c.Redirect(beego.URLFor("LoginController.Get"),302)
        c.StopRun()
    }
    username:=c.Ctx.GetCookie("username")
    fmt.Println(username)
}





