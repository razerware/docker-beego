package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
	_ "github.com/razerware/docker_beego/routers"
)

func main() {
	flag.Parse()
	models.Init()
	glog.Flush()
	beego.Run()
}
