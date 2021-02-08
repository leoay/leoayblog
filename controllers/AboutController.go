package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

//关于本站
func (c *AboutController) AboutWebSite() {
	c.Layout = "layouts/app.tpl"
	c.TplName = "pages/about_website.tpl"
}

//关于我
func (c *AboutController) AboutMe() {
	c.Layout = "layouts/app.tpl"
	c.TplName = "pages/about_me.tpl"
}
