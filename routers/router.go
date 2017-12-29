package routers

import (
	"docker-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/images",&controllers.ImagesController{})

    beego.Router("/home/:username",&controllers.HomeController{})
    beego.Router("/cluster_monitor",&controllers.ClusterController{},"*:ClusterMonitor")
    beego.Router("/cluster_detail",&controllers.ClusterController{},"*:ClusterDetail")
    beego.Router("/cluster_apply",&controllers.ClusterController{},"*:ClusterApply")

    beego.Router("/vm_detail",&controllers.VmController{},"*:VmDetail")
    beego.Router("/vm_detail_api",&controllers.VmController{},"*:VmDetailApi")
    beego.Router("/vm_apply",&controllers.VmController{},"*:VmApply")

    beego.Router("/service_apply",&controllers.ServiceController{},"*:ServiceApply")
    beego.Router("/service_detail",&controllers.ServiceController{},"*:ServiceDetail")
    beego.Router("/service_monitor",&controllers.ServiceController{},"*:ServiceMonitor")

	beego.Router("/containers",&controllers.ContainerController{})
	beego.Router("/lzy",&controllers.ContainerController{},"post:Lzy")
}
