//用户注册控制器
package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//只开放微信注册
//当微信扫码登陆后，检测到是新用户，那么就跳转到注册页面

//普通邮箱和用户名注册
func (c *RegisterController) Register() {
	//注册路由
	c.Ctx.WriteString("WSSDSS")
}
