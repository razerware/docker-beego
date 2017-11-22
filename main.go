package main

import (
	_ "docker-beego/routers"
	"github.com/astaxie/beego"
	"docker-beego/controllers"
	"fmt"
)

func main() {
	fmt.Println("sss")
	controllers.Init()
	beego.Run()
}
