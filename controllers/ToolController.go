package controllers

import (
	"github.com/astaxie/beego"
)

//ToolController 开发工具箱
type ToolController struct {
	beego.Controller
}

//WorkStation 工作空间
func (c *ToolController) WorkStation() {
	c.Layout = "layouts/app.html"
	c.TplName = "pages/workstation/workstation.html"
}
