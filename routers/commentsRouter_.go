package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "ClusterExpand",
			Router: `/cluster_expand`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("managerIp"),
				param.New("ip"),
				param.New("joinToken"),
				param.New("step"),
			),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "ClusterJoin",
			Router: `/cluster_join`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "ClusterList",
			Router: `/cluster_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "VmList",
			Router: `/vm_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "VmListAll",
			Router: `/vm_list_all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ServiceController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "ServiceApply",
			Router: `/service_apply`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:ServiceController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "ServiceApplyPost",
			Router: `/service_apply`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:TestController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:TestController"],
		beego.ControllerComments{
			Method: "MysqlQuery_Test",
			Router: `/test_mysql`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:TestController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:TestController"],
		beego.ControllerComments{
			Method: "ServiceApply_Test",
			Router: `/test_service`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["docker-beego/controllers:TestController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:TestController"],
		beego.ControllerComments{
			Method: "ServiceApply_Test2",
			Router: `/test_service`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
