package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"docker-beego/models"
	"github.com/golang/glog"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	if uid != nil {
		glog.V(1).Info("session exist")
		c.Redirect(beego.URLFor("HomeController.Get"), 302)
		c.StopRun()
	}
	glog.Info("need to login")
	c.TplName = "login.tpl"
}
func (c *LoginController) Post() {
	// 从Router拿到username
	var user User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	glog.Info(string(c.Ctx.Input.RequestBody))
	glog.Info(user)
	auth, uid := c.Auth(user)
	if auth == true {
		c.SetSession(user.Username, uid)
		c.Ctx.SetCookie("username", user.Username)
		url := beego.URLFor("HomeController.Get")
		glog.V(1).Info(url)
		c.Redirect(url, 302)
		c.StopRun()
	} else {
		c.Data["msg"] = "password or username error"
		c.TplName = "login.tpl"
	}

}

func (c *LoginController) Auth(user User) (bool, interface{}) {
	sql := "SELECT password,uid FROM user Where username='" + user.Username + "'"
	record := models.MysqlQuery(sql)
	fmt.Println(record)
	password, ok := record[0]["password"]
	if ok && password == user.Password {
		uid, _ := record[0]["uid"]
		return true, uid
	}
	return false, -1

}
