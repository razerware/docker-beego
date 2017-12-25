package routers

import (
	"docker-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/images",&controllers.ImagesController{})

    beego.Router("/hdr",&controllers.HdrController{})
    beego.Router("/test",&controllers.HdrController{},"*:Test")
    beego.Router("/test2",&controllers.HdrController{},"*:Test2")
    beego.Router("/cluster_apply",&controllers.HdrController{},"*:ClusterApply")
    beego.Router("/vm_detail",&controllers.HdrController{},"*:VmDetail")
    beego.Router("/vm_apply",&controllers.HdrController{},"*:VmApply")
    beego.Router("/service_apply",&controllers.HdrController{},"*:ServiceApply")
    beego.Router("/service_detail",&controllers.HdrController{},"*:ServiceDetail")
    beego.Router("/service_info",&controllers.HdrController{},"*:ServiceInfo")

	beego.Router("/containers",&controllers.ContainerController{})
	beego.Router("/lzy",&controllers.ContainerController{},"post:Lzy")
}
