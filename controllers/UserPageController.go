package controllers

//用户信息控制器
import "github.com/astaxie/beego"

//UserInfoController 用户信息控制器
type UserInfoController struct {
	AuthController
}

//Index 用户主页入口 GET
func (c *UserInfoController) Index() {

}

//UserBlogPub 获取当前用户的所有发布的文章信息 GET
func (c *UserInfoController) UserBlogPub() {

	//首先鉴权用户是否登陆
	if c.isLogin {
		userID := c.GetString(":userid")
		beego.Info("====================", userID)
		//判断登陆的用户与请求的用户是否一致
		if c.User.ID == c.User.ID {
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

//UserBlogDraf 获取当前用户的所有保存草稿的文章信息 GET
func (c *UserInfoController) UserBlogDraf() {

}
