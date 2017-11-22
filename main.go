package main

import (
	_ "docker-beego/routers"
	"github.com/astaxie/beego"
	"docker-beego/controllers"
	"fmt"
)

func main() {
	controllers.Init()
	beego.Run()
}
