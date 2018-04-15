package routers

import (
	"docker-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/images",&controllers.ImagesController{})

    beego.Router("/home/?:username",&controllers.HomeController{})
	//beego.Router("/cluster_detail",&controllers.ClusterController{},"*:ClusterDetail")
    //beego.Router("/list_cluster",&controllers.ClusterController{},"*:ListCluster")
    //beego.Router("/cluster_apply",&controllers.ClusterController{},"*:ClusterApply")

    //beego.Router("/vm_detail",&controllers.VmController{},"*:VmDetail")
    //beego.Router("/vm_detail_api",&controllers.VmController{},"*:VmDetailApi")
    //beego.Router("/vm_apply",&controllers.VmController{},"*:VmApply")

    //beego.Router("/service_apply",&controllers.ServiceController{},"*:ServiceApply")
    //beego.Router("/service_detail",&controllers.ServiceController{},"*:ServiceDetail")
	//beego.Router("/list_service",&controllers.ServiceController{},"*:ListService")
    //beego.Router("/service_monitor",&controllers.ServiceController{},"*:ServiceMonitor")

    beego.Router("/login", &controllers.LoginController{},"*:LogIn")
	beego.Router("/logout", &controllers.LoginController{},"*:LogOut")
	beego.Router("/containers",&controllers.ContainerController{})
	beego.Router("/lzy",&controllers.ContainerController{},"post:Lzy")

	beego.Include(&controllers.ClusterController{})
	beego.Include(&controllers.TestController{})
	beego.Include(&controllers.VmController{})
	beego.Include(&controllers.ServiceController{})
}
