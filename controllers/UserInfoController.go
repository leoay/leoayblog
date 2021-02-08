package controllers

import "github.com/astaxie/beego"

type UserInfoController struct {
	UserStatusController
}

//获取当前用户的所有发布的文章信息
func (c *UserInfoController) UserBlogPub() {

	//首先鉴权用户是否登陆
	if c.isLogin {
		user_id := c.GetString(":userid")
		beego.Info("====================", user_id)
		//判断登陆的用户与请求的用户是否一致
		if c.User.Id == c.User.Id {
			c.Ctx.WriteString("OkUser")

		} else {
			//跳转到错误页
			c.Ctx.WriteString("WrongPage")
		}

		c.Ctx.WriteString("Hello")
	} else {
		//如果session超时，则显示登陆按钮
		c.Redirect("/", 302)
	}
}

//获取当前用户的所有保存草稿的文章信息
func (c *UserInfoController) UserBlogDraf() {

}
