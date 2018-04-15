package controllers

import (
	"github.com/astaxie/beego"
	"docker-beego/models"
	"github.com/golang/glog"
	"strings"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) LogIn() {
	//uid:=c.GetUid()
	//if uid <0 {
	//	glog.V(1).Info("session exist")
	//	c.Redirect(beego.URLFor("HomeController.Get"), 302)
	//	c.StopRun()
	//}
	//glog.Info("need to login")
	if c.userId >= 0 {
		c.redirect(beego.URLFor("HomeController.Get"))
	}
	beego.ReadFromRequest(&c.Controller)
	if c.Ctx.Request.Method == "POST" {
		var user models.User
		//json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		//glog.Info(string(c.Ctx.Input.RequestBody))
		//glog.Info(user)
		user.Username = strings.TrimSpace(c.GetString("username"))
		user.Password = strings.TrimSpace(c.GetString("password"))
		glog.Info(user)
		flash := beego.NewFlash()
		auth, uid := c.Auth(user)
		if auth == true {
			c.SetSession(user.Username, uid)
			c.Ctx.SetCookie("username", user.Username)
			glog.Info("auth true and cookie set")

			v, ok := uid.(int)
			if ok && v == 1 {
				c.Ctx.SetCookie("url", "10.109.252.172:8888")
			} else {
				c.Ctx.SetCookie("url", "10.109.252.172:10000")
			}
			c.redirect(beego.URLFor("HomeController.Get"))
		} else {
			glog.Info("auth error")
			flash.Error("账号或密码错误")
			flash.Store(&c.Controller)
			c.redirect(beego.URLFor("LoginController.LogIn"))
		}
	}
	c.TplName = "login.tpl"
}

//func (c *LoginController) Post() {
//	// 从Router拿到username
//	var user models.User
//	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
//	glog.Info(string(c.Ctx.Input.RequestBody))
//	glog.Info(user)
//	auth, uid := c.Auth(user)
//	if auth == true {
//		c.SetSession(user.Username, uid)
//		c.Ctx.SetCookie("username", user.Username)
//		url := beego.URLFor("HomeController.Get")
//		glog.V(1).Info(url)
//		c.Redirect(url, 302)
//		c.StopRun()
//	} else {
//		c.Data["msg"] = "password or username error"
//		c.TplName = "login.tpl"
//	}
//
//}

func (c *LoginController) Auth(user models.User) (bool, interface{}) {
	if user.Username == "" {
		return false, -1
	}
	sql := "SELECT password,uid FROM user Where username='" + user.Username + "'"
	record := models.MysqlQuery(sql)
	glog.Info(record)
	if record!=nil{
		password, ok := record[0]["password"]
		if ok && password == user.Password {
			uid, _ := record[0]["uid"]
			return true, uid
		}
	}
	glog.Info("Auth error")
	return false, -1
}

func (c *LoginController) LogOut() {
	username := c.Ctx.GetCookie("username")
	c.DelSession(username)
	c.Ctx.SetCookie("username", "")
	c.redirect(beego.URLFor("LoginController.LogIn"))
}
