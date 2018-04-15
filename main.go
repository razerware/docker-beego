package main

import (
	_"docker-beego/routers"
	"github.com/astaxie/beego"
	"flag"
	"github.com/golang/glog"
	"docker-beego/models"
)

func main() {
	flag.Parse()
	models.Init()
	glog.Flush()
	beego.Run()
}
