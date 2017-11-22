package main

import (
	_ "docker-beego/routers"
	"github.com/astaxie/beego"
	"docker-beego/controllers"
)

func main() {
	controllers.Init()
	beego.Run()
}
