package controllers

//LogoutController 登出结构体
type LogoutController struct {
	AuthController
}

//Logout For Logout
func (c *LogoutController) Logout() {
	//Todo: 如果当前账户已登陆，则跳转到首页
	if c.isLogin {
		c.DelSession("userid")
		c.DelSession("password")
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.SetCookie("status", "success", 1)
		c.Ctx.SetCookie("message", "登出成功，再见！", 1)
		c.Ctx.Redirect(302, "/")
	} else {
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	}
}
