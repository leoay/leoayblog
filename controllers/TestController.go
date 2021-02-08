package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

//登陆界面
func (c *TestController) Index() {
	c.TplName = "pages/test/index.tpl"
}
