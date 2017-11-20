package routers

import (
	"docker-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/images",&controllers.ImagesController{})
	beego.Router("/containers",&controllers.ContainerController{})
}
