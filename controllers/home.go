package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	// 从Router拿到username
	//username:=c.Ctx.GetCookie("username")
	//glog.V(1).Info(username)
	//uid := c.GetSession(username)
	//glog.V(1).Info(uid)
	//if uid==nil{
	//    glog.Info("session not exist")
	//    c.Redirect(beego.URLFor("LoginController.Get"),302)
	//    c.StopRun()
	//}
	c.TplName = "layout.tpl"
}
