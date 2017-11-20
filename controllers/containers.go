package controllers

import (
"github.com/astaxie/beego"
	"fmt"
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
}
func (c *ContainerController) GetContainers() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}