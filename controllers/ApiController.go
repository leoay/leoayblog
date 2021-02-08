package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

//生成校验码，用于在微信端调用
func (c *ApiController) GenCheckCodePre() {
	c.TplName = "pages/getCode.tpl"
	//c.Ctx.WriteString("AAAAA")
	//c.Ctx.Redirect(302, "http://baidu.com")
}

//生成校验码，用于在微信端调用
func (c *ApiController) GenCheckCodeOk() {
	c.Ctx.WriteString("AAAAA")
}

//校验验证码，用于校验从微信端获取的验证码是否有效
func (c *ApiController) CheckCode() {
	//校验通过后，将浏览器的cookies 进行持久化存储
}
