package controllers

import (
	"github.com/razerware/docker_beego/models"
	"fmt"
	"github.com/golang/glog"
	"encoding/json"
)

type TestController struct {
	BaseController
}

// @router /test_mysql [get]
func (c *TestController) MysqlQuery_Test() {
	//uid := 1
	sql := fmt.Sprintf("SELECT ip FROM `vm_info` WHERE swarm_id=''")
	c.CustomAbort(203, "ggg")
	////swarm_id:=1
	////sql := fmt.Sprintf("SELECT * FROM `service` WHERE swarm_id=%d", swarm_id)
	////temp:=ClusterInfo{}
	//record := models.MysqlQuery(sql, ClusterInfo{})
	//fmt.Println(record)
	record := models.MysqlQuery(sql)
	if len(record) < 1 {
		glog.Error("no result")
		return
	}
	//glog.V(1).Info(record)
	//temp, err := json.Marshal(record)
	//if err != nil {
	//	glog.Error(err)
	//}
	////st:=ClusterInfo{}
	////err = json.Unmarshal(temp, &st)
	////if err != nil {
	////	glog.Error(err)
	////}
	c.Data["json"] = record
	c.ServeJSON()
}

// @router /test_service [get]
func (c *TestController) ServiceApply_Test() {
	//js:=`{"Name": "whoami_test", "Image": "10.109.252.163:5000/emilevauge/whoami",
	//"Constraints":"node.role==worker",
	//"Target": "traefik-net",
	//"Replicas": 4,
	//"traefik.port": "80",
	//"traefik.frontend.rule":"Host:test.com"}`
	//glog.Info("123")
	//MyPost("http://127.0.0.1:8080/service_apply_post",[]byte(js))
	//glog.Info("456")
	c.Data["json"] = "gg"
	c.ServeJSON()
}

// @router /test_service [post]
func (c *TestController) ServiceApply_Test2() {
	//js:=`{"Name": "whoami_test", "Image": "10.109.252.163:5000/emilevauge/whoami",
	//"Constraints":"node.role==worker",
	//"Target": "traefik-net",
	//"Replicas": 4,
	//"traefik.port": "80",
	//"traefik.frontend.rule":"Host:test.com"}`
	//glog.Info("123")
	//MyPost("http://127.0.0.1:8080/service_apply_post",[]byte(js))
	//glog.Info("456")

	var fss models.FrontendSS
	json.Unmarshal(c.Ctx.Input.RequestBody, &fss)
	sql := fmt.Sprintf("INSERT INTO `service` (`name`,`upper_limit`, `lower_limit`, `step`, "+
		"`cpu_lower`, `cpu_upper`, `mem_lower`, `mem_upper`) "+
		"VALUES ('%s','%d', '%d', '%d', '%d', '%d', '%d', '%d')",
		fss.Name, fss.UpperLimit, fss.LowerLimit, fss.Step, fss.CpuLower, fss.CpuUpper, fss.MemLower, fss.MemUpper)
	fmt.Println(sql)
	last, row, err := models.MysqlInsert(sql)
	glog.Info(last, row, err)
	c.Data["json"] = "gg"
	c.ServeJSON()

}

// @router /test_session [get]
func (c *TestController) Session_test() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	c.Data["json"] = uid
	c.ServeJSON()

}

// @router /test_uid [*]
func (c *TestController) Uid_test() {
	c.Data["json"]=c.userId
	c.ServeJSON()
}