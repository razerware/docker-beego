package controllers

import (
"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

type ContainerController struct {
	beego.Controller
}
type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

type user2 struct {
	Id    int `json:"-"`
	Username  interface{} `json:"username"`
	Age   int         `json:"age"`
	Text string      `json:"text"`
}
func (c *ContainerController) Get() {
	c.TplName = "test.tpl"
}
func (this *ContainerController) Post() {
	u := user{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		fmt.Println(u.Id)
	}
	this.Data["Website"] = u.Name
	this.TplName="index.tpl"
	var ob user2
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	fmt.Println(string(this.Ctx.Input.RequestBody))
	this.Data["json"] = ob
	this.ServeJSON()
}
func (c *ContainerController) GetContainers() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}