package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"io/ioutil"
	"docker-beego/client"
)

type ImagesController struct {
	beego.Controller
}

func (c *ImagesController) Get() {
	c.TplName = "form.tpl"
}
func (c *ImagesController) Post() {
	c.TplName = "test.tpl"
}

func (c *ImagesController) PullImages(name string ,tag string){
	url:=client.GetClient()
	req, _ := http.NewRequest("POST", url,nil)
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	q:=req.URL.Query()

	q.Add("fromImage","hello-world")
	q.Add("tag","latest")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}