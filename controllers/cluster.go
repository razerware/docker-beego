package controllers

import (
	"github.com/astaxie/beego"
	"docker-beego/elbController"
	"fmt"
	"docker-beego/models"
)

type ClusterController struct {
	beego.Controller
}

type ClusterInfo struct {
	ID         int
	Token      string `json:"token"`
	User
	elbController.ElasticInfo
	manager_ip string
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

func (c *ServiceController) ListCluster() {
	uid := c.Ctx.GetCookie("session_id")
	var manager_ip string
	var manager_ips []string
	db := models.DB
	if uid != "" {
		rows, err := db.Query("SELECT manager_ip FROM cluster_info WHERE uid=" + uid)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rows)
			for rows.Next() {
				if err = rows.Scan(&manager_ip); err != nil {
					fmt.Println(err)
					c.Data["json"] = err
					c.ServeJSON()
				}
			}
		}
		fmt.Println(manager_ip)
		manager_ips = append(manager_ips, manager_ip)
		c.Data["json"] = manager_ips
		c.ServeJSON()
	}

}

func (c *ClusterController) Test() {
	c.TplName = "cluster_apply.tpl"
}
