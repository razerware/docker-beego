package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type BaseController struct {
	beego.Controller
	userId int
}

func (c *BaseController) Prepare() {
	c.userId = -1
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	glog.V(1).Info(username, ",", uid)
	v, ok := uid.(int)
	if ok {
		c.userId = v
		glog.V(1).Info("session exist: username: ", username, ",uid: ", uid)
	} else {
		glog.Info("session not exist")
		controllerName, actionName := c.GetControllerAndAction()
		//检查是否正在路由在login页面
		if c.userId == -1 && (controllerName != "LoginController" && actionName != "logIn") {
			glog.V(1).Info("Not in Login")
			c.redirect(beego.URLFor("LoginController.LogIn"))
		} else {
			glog.V(1).Info("This is in Login")
		}
	}
}

//
//func (c *BaseController) GetUid() int {
//	username := c.Ctx.GetCookie("username")
//	uid := c.GetSession(username)
//	v, ok := uid.(int)
//	if ok {
//		return v
//	}
//	return -1
//
//}

func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}
