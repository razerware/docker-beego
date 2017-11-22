package controllers

import "docker-beego/models"

func Init(){
	models.MysqlConnect()
}
