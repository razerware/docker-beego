package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["docker-beego/controllers:TestController"] = append(beego.GlobalControllerRouter["docker-beego/controllers:TestController"],
		beego.ControllerComments{
			Method: "MysqlQuery_Test",
			Router: `/test_mysql`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
