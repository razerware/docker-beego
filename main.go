package main

import (
	_ "docker-beego/routers"
	"github.com/astaxie/beego"
	"flag"
	"github.com/golang/glog"
	"docker-beego/controllers"
)

func main() {
	flag.Parse()
	controllers.Init()
	glog.Flush()
	beego.Run()
}
