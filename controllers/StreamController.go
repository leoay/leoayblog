package controllers

import (
	"github.com/astaxie/beego"
)

//LoginController 控制器结构
type StreamController struct {
	beego.Controller
}

//Index 登陆界面
func (c *StreamController) Show() {
	c.Layout = "layouts/app.html"
	c.TplName = "pages/showvideo.html"
}
