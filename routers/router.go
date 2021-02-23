package routers

import (
	"WriterReaderLog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//上传图片路由
	beego.Router("/api/upload_img", &controllers.UploadImgController{})
	//上传文件路由

	//登陆认证模块
	//登陆界面路由
	beego.Router("/user/login", &controllers.LoginController{}, "get:Index")
	beego.Router("/user/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/user/register", &controllers.RegisterController{}, "post:Register")
	beego.Router("/user/logout", &controllers.LogoutController{}, "get:Logout")

	//代码技术
	//技术专栏
	beego.Router("/codeTech/tecCol", &controllers.CodeTechController{}, "get:TechColumn")

	//####写作工具和文章管理#######################
	beego.Router("/writeTool", &controllers.WriteToolController{}, "get:Index")
	beego.Router("/writeTool/save_temp", &controllers.WriteToolController{}, "post:SaveTemp")
	beego.Router("/writeTool/publish", &controllers.WriteToolController{}, "post:Publish")
	beego.Router("/writeTool/:id/edit_article", &controllers.WriteToolController{}, "get:EditArticle")
	beego.Router("/writeTool/delete_article", &controllers.WriteToolController{}, "post:DeleteArticle")

	//文章细节展示工具
	beego.Router("/show_article/:id", &controllers.WriteToolController{}, "get:ShowArticleDetail")

	//#####################################
	//碎碎念
	beego.Router("/creativeAbility/brokenThoughts", &controllers.CreatorController{}, "get:DaliyThinkIndex")
	//公众号文章
	beego.Router("/creativeAbility/wePubWriting", &controllers.CreatorController{}, "get:WeArticleIndex")

	//关于路由
	beego.Router("/about/leoaylab", &controllers.AboutController{}, "get:AboutWebSite")
	beego.Router("/about/leoay", &controllers.AboutController{}, "get:AboutMe")

	//图片池上传
	beego.Router("/picpool/upload", &controllers.PicPoolController{}, "post:Upload")
	//获取所有图片信息
	beego.Router("/picpool/queryimages", &controllers.PicPoolController{}, "get:QueryImages")

	//github oauth认证路由
	beego.Router("/api/github_oauth", &controllers.LoginController{}, "get:Github_oauth")

	//公众号关注API测试
	beego.Router("/api/getWeCode", &controllers.ApiController{}, "get:GenCheckCodePre")

	//视频播放器
	beego.Router("/api/play", &controllers.MainController{}, "get:VideoPlayer")

	//PDF reader
	beego.Router("/reader/pdf", &controllers.MainController{}, "get:PdfReader")

	//用户信息
	//用户的博客显示页
	beego.Router("/user/blog/:userid", &controllers.UserInfoController{}, "get:UserBlogPub")

}
