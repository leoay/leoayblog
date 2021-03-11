package controllers

import (
	"github.com/astaxie/beego"
)

type ToolController struct {
	beego.Controller
}

//WorkStation 工作空间
func (c *ToolController) WorkStation() {
	c.Layout = "layouts/app.tpl"
	c.TplName = "pages/workstation/workstation.tpl"
}
