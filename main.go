package main

import (
	_"github.com/razerware/docker_beego/routers"
	"github.com/astaxie/beego"
	"flag"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
)

func main() {
	flag.Parse()
	models.Init()
	glog.Flush()
	beego.Run()
}
