package controllers

type ImagesController struct {
	BaseController
}

func (c *ImagesController) Get() {
	c.TplName = "form.tpl"
}
func (c *ImagesController) Post() {
	c.TplName = "index.tpl"
}
