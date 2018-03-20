package controllers

import (
"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"docker-beego/models"
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
	Name  interface{} `json:"username"`
	Age   int         `json:"age"`
	Email string      `json:"email"`
}
func (c *ContainerController) Get() {
	if models.DbError!=nil{
		fmt.Println(models.DbError)
		c.TplName = "test.tpl"
	}
	db:=models.DB
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		c.TplName = "test.tpl"
		return
	}else{
		fmt.Println("ok")
	}
	rows,err:=db.Query("SELECT * FROM service")
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("ok rows")
	}
	for rows.Next()  {
		var name string
		if err=rows.Scan(&name);err!=nil{
			fmt.Println(err)
			fmt.Println("ggggggg")
		}else{
			fmt.Println(name)
			fmt.Println("ssssss")
		}
	}
	c.TplName = "test.tpl"
}

func (this *ContainerController) Lzy() {
	var ob user2
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	this.Data["json"] = ob
	this.ServeJSON()
}

func (this *ContainerController) Post() {
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