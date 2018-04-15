package controllers

import (
	"fmt"
	"docker-beego/models"
)

type VmController struct {
	BaseController
}

// @router /vm_detail [get]
func (c *VmController) VmDetail() {
	c.TplName = "vm_detail.tpl"
}

// @router /vm_list_all [get]
func (c *ClusterController) VmListAll() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	//var uid int
	//c.Ctx.Input.Bind(&uid, "uid")
	sql := fmt.Sprintf("SELECT * FROM vm_info WHERE uid=%d", uid)
	record := models.MysqlQuery(sql)
	fj := models.FrontendJson{0, "", len(record), record}
	c.Data["json"] = fj
	c.ServeJSON()
}

// @router /vm_list [get]
func (c *ClusterController) VmList() {
	//username := c.Ctx.GetCookie("username")
	//uid := c.GetSession(username)
	var swarmId string
	c.Ctx.Input.Bind(&swarmId, "swarmId")
	sql := fmt.Sprintf("SELECT * FROM `vm_info` WHERE swarm_id ='%s'", swarmId)
	record := models.MysqlQuery(sql)
	fj := models.FrontendJson{0, "", len(record), record}
	c.Data["json"] = fj
	c.ServeJSON()
}
