package controllers

import (
	"github.com/astaxie/beego"
	"docker-beego/elbController"
	"fmt"
	"docker-beego/models"
	"encoding/json"
	"github.com/golang/glog"
)

type ClusterController struct {
	beego.Controller
}

type ClusterInfo []struct {
	Swarm_ID   int    `json:"swarm_id"`
	Token      string `json:"token"`
	Manager_ip string `json:"manager_ip"`
	User
	elbController.ElasticInfo
}

type ClusterInit struct {
	//ListenAddr      string `json:"ListenAddr"`
	AdvertiseAddr string `json:"AdvertiseAddr"`
	Spec struct {
		Name string `json:"Name"`
	} `json:"Spec"`
}

type ClusterJoin struct {
	ListenAddr    string `json:"listenAddr"`
	AdvertiseAddr string   `json:"AdvertiseAddr"`
	JoinToken     string   `json:"JoinToken"`
	RemoteAddrs   []string `json:"RemoteAddrs"`
}

type FrontendCI struct {
	SwarmId      string `json:"SwarmId"`
	AdvertiseAddr string `json:"AdvertiseAddr"`
	Name          string `json:"Name"`
	elbController.ElasticInfo
}

type FrontendCJ struct {
	//ListenAddr    string `json:"listenAddr"`
	AdvertiseAddr string   `json:"AdvertiseAddr"`
	JoinToken     string   `json:"JoinToken"`
	ManagerIp   string `json:"ManagerIp"`
}


type FrontendJson struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int `json:"count"`
	Data interface{} `json:"data"`
}

func (c *ClusterController) ClusterMonitor() {
	c.TplName = "cluster_monitor.tpl"
}

func (c *ClusterController) ClusterDetail() {
	c.TplName = "cluster_detail.tpl"
}

func (c *ClusterController) ClusterApply() {
	c.TplName = "cluster_apply.tpl"
}

func (c *ClusterController) ClusterInit() {
	//username := c.Ctx.GetCookie("username")
	//uid := c.GetSession(username)
	uid := 1
	var manager_ip, name string
	c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	c.Ctx.Input.Bind(&name, "name")
	var fci FrontendCI
	json.Unmarshal(c.Ctx.Input.RequestBody, &fci)
	glog.Info(fci)
	var ci ClusterInit
	ci.AdvertiseAddr = manager_ip
	ci.Spec.Name = name

	send_ss, ok := json.Marshal(ci)
	if ok != nil {
		glog.Error(ok)
	}
	url := "http://" + manager_ip + ":2375/swarm/init"
	code, body, err := MyPost(url, send_ss)
	if err != nil || code > 200 {
		glog.Error(code)
		glog.Error(err)
	}
	glog.Info(string(send_ss))
	glog.Info(body)
	fci.SwarmId = string(body)
	RegistCluster(fci, uid)

	c.Data["json"] = ci
	c.ServeJSON()
}
// @router /cluster_join [post]
func (c *ClusterController) ClusterJoin() {

	var fcj FrontendCJ
	json.Unmarshal(c.Ctx.Input.RequestBody, &fcj)
	var cj ClusterJoin
	cj.ListenAddr="0.0.0.0:2377"
	cj.AdvertiseAddr = fcj.AdvertiseAddr
	cj.JoinToken = fcj.JoinToken
	cj.RemoteAddrs = append(cj.RemoteAddrs, fcj.ManagerIp+":2377")
	sendCJ, ok := json.Marshal(cj)
	if ok!=nil{
		glog.Error(ok)
		c.Data["json"] = ""
		c.ServeJSON()
	}
	glog.Info(string(sendCJ), ok)
	url:="http://"+fcj.AdvertiseAddr+":2375/swarm/join"
	glog.Info(url)
	code,body,err:=MyPost(url,sendCJ)
	if err!=nil||code>200{
		glog.Error("swarm join failed",code)
		glog.Error(err)
		c.Data["json"] = fcj
		c.ServeJSON()
	}
	glog.Info("swarm join successed ",string(sendCJ))
	glog.Info("resp is ",string(body))
	c.Data["json"] = fcj
	c.ServeJSON()
}


// @router /cluster_list [get]
func (c *ClusterController) ClusterList() {
	username := c.Ctx.GetCookie("username")
	uid := c.GetSession(username)
	//var uid int
	//c.Ctx.Input.Bind(&uid, "uid")
	sql := fmt.Sprintf("SELECT * FROM cluster_info WHERE uid=%d", uid)
	record := models.MysqlQuery(sql)
	fj:=FrontendJson{0,"",len(record),record}
	c.Data["json"] = fj
	c.ServeJSON()
}





func RegistCluster(fci FrontendCI, uid int) {

	//c.Ctx.Input.Bind(&uid, "uid")
	//json.Unmarshal(c.Ctx.Input.RequestBody, &fci)
	sql := fmt.Sprintf("INSERT INTO `cluster_info` (`swarm_id`,`name`,`uid`,`manager_ip`,"+
		"`upper_limit`, `lower_limit`, `step`,`cpu_lower`, `cpu_upper`, `mem_lower`, `mem_upper`) "+
		"VALUES ('%s','%s','%d','%s','%d', '%d', '%d', '%d', '%d', '%d', '%d')",
		fci.SwarmId, fci.Name, uid, fci.AdvertiseAddr,
		fci.UpperLimit, fci.LowerLimit, fci.Step, fci.CpuLower, fci.CpuUpper, fci.MemLower, fci.MemUpper)
	fmt.Println(sql)
	last, row, err := models.MysqlInsert(sql)
	glog.Info(last, row, err)

}
// @router /cluster_expand [post]
func (c *ClusterController)ClusterExpand(managerIp string, ip string, joinToken string,step int) (bool, error) {
	//var manager_ip,name string
	//c.Ctx.Input.Bind(&manager_ip, "manager_ip")
	sql:="SELECT ip FROM `vm_info` WHERE `uid` = 1 AND `swarm_id` = ''"
	record:=models.MysqlQuery(sql)
	if len(record)<=step{
		for _,v:=range record{
			c.ClusterJoin()
			fmt.Println(v)
		}
	}

	return false,nil
}


func (c *ClusterController) Test() {
	c.TplName = "cluster_apply.tpl"
}
