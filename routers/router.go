package routers

import (
	"docker-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/images",&controllers.ImagesController{})

    beego.Router("/home",&controllers.HdrController{})
    beego.Router("/cluster_monitor",&controllers.HdrController{},"*:ClusterMonitor")
    beego.Router("/cluster_detail",&controllers.HdrController{},"*:ClusterDetail")
    beego.Router("/cluster_apply",&controllers.HdrController{},"*:ClusterApply")
    beego.Router("/vm_detail",&controllers.HdrController{},"*:VmDetail")
    beego.Router("/vm_detail_api",&controllers.HdrController{},"*:VmDetailApi")
    beego.Router("/vm_apply",&controllers.HdrController{},"*:VmApply")
    beego.Router("/service_apply",&controllers.HdrController{},"*:ServiceApply")
    beego.Router("/service_detail",&controllers.HdrController{},"*:ServiceDetail")
    beego.Router("/service_monitor",&controllers.HdrController{},"*:ServiceMonitor")

	beego.Router("/containers",&controllers.ContainerController{})
	beego.Router("/lzy",&controllers.ContainerController{},"post:Lzy")
}
