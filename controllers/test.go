package controllers

import (
	"github.com/astaxie/beego"
	"docker-beego/models"
	"fmt"
	"encoding/json"
)

type TestController struct {
	beego.Controller
}

// @router /test_mysql [get]
func (c *TestController) MysqlQuery_Test() {
	record := models.MysqlQuery("SELECT * FROM `service`")
	fmt.Println(record)
	temp, err := json.Marshal(record)
	if err == nil {
		fmt.Println(string(temp))
	}
	temp2:=LBspec{}
	err=json.Unmarshal(temp,&temp2)
	fmt.Println(err)
	c.Data["json"] = temp2
	c.ServeJSON()
}
