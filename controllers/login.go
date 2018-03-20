package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"docker-beego/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){
	//url:=beego.URLFor("ClusterController.ClusterApply")
	session_id:=c.Ctx.GetCookie("session_id")
	v := c.GetSession(session_id)
	if v!=nil{
		fmt.Println("session exist")
		c.Redirect(beego.URLFor("HomeController.Get"),302)
		c.StopRun()
	}
	fmt.Println("login")
	c.TplName = "login.tpl"
	//fmt.Println(url)
	//c.Redirect(url,302)
	//c.StopRun()
}
func (c *LoginController) Post() {
	// 从Router拿到username
	var user User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	fmt.Println(string(c.Ctx.Input.RequestBody))
	fmt.Println(user)
	auth,msg:=c.Auth(user)
	if auth==true{
		c.SetSession(msg,user.Username)
		c.Ctx.SetCookie("session_id",msg)
		url := beego.URLFor("HomeController.Get")
		fmt.Println(url)
		c.Redirect(url, 302)
		c.StopRun()
	} else {
		c.Data["msg"] = msg
		c.TplName = "error.tpl"
	}
	//v := c.GetSession(user.Username)

	//if v != nil {
	//	c.Data["num"] = 0
	//	url:=beego.URLFor("HomeController.Get")
	//	fmt.Println(url)
	//	c.Redirect(url,302)
	//	c.StopRun()
	//}else {
	//
	//	fmt.Println(auth,msg)
	//
	//}
	//fmt.Println(v)
	//c.Data["json"] = c.GetSession(user.Username)
	//c.ServeJSON()

}
func (c *LoginController) Test() {
	// 从Router拿到username
	v := c.GetSession("lzy")
	c.Data["json"]=v
	c.ServeJSON()
}

func (c *LoginController) Auth(user User) (bool, string){
	if models.DbError!=nil{
		fmt.Println(models.DbError)
	}
	db:=models.DB
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return false,fmt.Sprint(err)
	}else{
		fmt.Println("ok")
	}
	rows,err:=db.Query("SELECT password,uid FROM user Where username='"+user.Username+"'")
	if err!=nil{
		fmt.Println(err)
	}else{
		for rows.Next()  {
			var password string
			var uid string
			if err=rows.Scan(&password,&uid);err!=nil{
				fmt.Println(err)
				return false,fmt.Sprint(err)
			}else if password==user.Password{
				fmt.Println(password,uid)
				return true,uid
			}else {
				return false,"wrong password"
			}
		}
	}
	return false,fmt.Sprint(err)
}

