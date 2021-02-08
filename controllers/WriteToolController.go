package controllers

//写作工具控制器

import (
	"WriterReaderLog/models"
	"WriterReaderLog/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WriteToolController struct {
	UserStatusController
}

//写作工具首页
func (c *WriteToolController) Index() {
	if c.isLogin {
		beego.Info("已登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		c.Data["newarticle"] = 1
		c.Layout = "layouts/app.tpl"
		c.TplName = "pages/writeTool/all_writer.tpl"
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	}
}

//保存草稿
func (c *WriteToolController) SaveTemp() {
	//获取文章编辑表单
	markdownText := c.GetString("leoaylog-editor-markdown-doc")
	//碎碎念的文章标题格式固定为xx年xx月xx日——碎碎念
	title := c.GetString("title")
	//一级标题
	firstClass := c.GetString("article_class1")
	//二级标题
	secondClass := c.GetString("article_class2")

	htmlCode := c.GetString("leoaylog-editor-html-code")
	beego.Info("一级分类", firstClass)
	beego.Info("二级分类", secondClass)
	beego.Info("标题", title)
	beego.Info("markdown", markdownText)
	beego.Info("htmlCode", htmlCode)

	dsn := "root:111111@tcp(127.0.0.1:3306)/leoaylab?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Err is")
	}

	//db.Model(&models.Article_bases{}).Where("id = ?", 2).Update("author", "klkdlsajklsajd")
	utils.DB.Model(&models.Article_bases{}).Where("id = ?", 2).Update("author", "cccccccccccccccccccccccccccccc")

	//暂时将markdown写到文件,并加上时间戳
	// fileName := "articles/" + strings.TrimSpace(title) + "#*#*#" + strconv.FormatInt(time.Now().Unix(), 10) + ".md"
	// dstFile, err := os.Create(fileName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// defer dstFile.Close()
	// dstFile.WriteString(markdownText + "\n")

	var article_class models.Article_classes
	db.Where("first_class = ? AND second_class = ?", firstClass, secondClass).Find(&article_class)
	istemp := true
	beego.Info("first_class", firstClass)
	beego.Info("second_class", secondClass)
	beego.Info("Id", article_class.Id)
	//db.Select("name", "age").Find(&users)

	article_base := models.Article_bases{Title: title, ClassId: article_class.Id, UserId: c.User.Id, Content: htmlCode, IsTemp: istemp, Content_md: markdownText}
	db.Create(&article_base) // 通过数据的指针来创建
	//保存草稿后，返回到编辑页面
	c.Ctx.SetCookie("status", "success", 1)
	c.Ctx.SetCookie("message", "恭喜！保存草稿成功！", 1)
	c.Ctx.Redirect(302, "/")
}

//发布文章
func (c *WriteToolController) Publish() {
	//登录之后才进行保存
	if c.isLogin {
		newarticle := c.GetString("newarticle")
		//获取文章编辑表单
		markdownText := c.GetString("leoaylog-editor-markdown-doc")
		title := c.GetString("title")
		firstClass := c.GetString("article_class1")
		secondClass := c.GetString("article_class2")
		summary := c.GetString("summary")
		mainPicPath := c.GetString("main_pic_path")
		htmlCode := c.GetString("leoaylog-editor-html-code")
		var article_class models.Article_classes
		utils.DB.Where("first_class = ? AND second_class = ?", firstClass, secondClass).Find(&article_class)
		istemp := false
		if newarticle == "1" {
			utils.DB.Create(&models.Article_bases{Title: title, Author: c.User.Name, EditorId: 1, MainPicPath: mainPicPath, Summary: summary, UserId: c.User.Id, ClassId: article_class.Id, Content: htmlCode, ApprovalNum: 10, StarNum: 10, IsTemp: istemp, Content_md: markdownText, Comments: "leoay", CreatedAt: uint64(time.Now().Unix()), UpdatedAt: uint64(time.Now().Unix())}) // 通过数据的指针来创建
		} else {
			ariid, _ := strconv.Atoi(c.GetString("articleid"))
			utils.DB.Model(&models.Article_bases{Id: uint(ariid)}).Updates(models.Article_bases{Title: title, ClassId: article_class.Id, MainPicPath: mainPicPath, Summary: summary, Content: htmlCode, ApprovalNum: 10, StarNum: 10, IsTemp: istemp, Content_md: markdownText, Comments: "leoay", UpdatedAt: uint64(time.Now().Unix())})
		}
	}
	c.Ctx.SetCookie("status", "success", 1)
	c.Ctx.SetCookie("message", "恭喜！你已成功发布一篇文章，快写下一篇吧！", 1)
	c.Ctx.Redirect(302, "/")
}

//重新编辑文章
func (c *WriteToolController) EditArticle() {
	artId := c.GetString(":id")
	if c.isLogin {
		var articles models.Article_bases
		utils.DB.Where("id=?", artId).Find(&articles)
		beego.Info("一登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		c.Data["markdown"] = articles.Content_md
		c.Data["artTitle"] = articles.Title
		c.Data["articleid"] = articles.Id
		c.Data["newarticle"] = 0
		classid := articles.ClassId
		var classT models.Article_classes
		utils.DB.Where("id=?", classid).Find(&classT)
		c.Data["firstClass"] = classT.FirstClass
		c.Data["secondClass"] = classT.SecondClass

		c.Layout = "layouts/app.tpl"
		c.TplName = "pages/writeTool/all_writer.tpl"
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	}
}

type ArticleInfo struct {
	Id                      int
	ArticleTitle            string
	ArticleContent          string
	ArticleFirstClass       string
	ArticleSecondClass      string
	ArticleAuthor           string
	ArticleCreateTimeSub    string
	ArticleUpdateTimeSub    string
	ArticleCreateTimeUTC    string
	ArticleUpdateTimeUTC    string
	ArticleFirstClassRoute  string
	ArticleSecondClassRoute string
}

//计算时间
func TimeTrans(timestamp uint64) (string, string) {
	timeSubUnix := uint64(time.Now().Unix()) - timestamp
	beego.Info(timeSubUnix)
	//years := math.Floor(float64(timeSubUnix / (24 * 3600)))
	//var days = Math.floor(dateDiff / (24 * 3600 * 1000));//计算出相差天数
	days := math.Floor(float64(timeSubUnix / (24 * 3600)))

	leave1 := math.Floor(float64(timeSubUnix % (24 * 3600))) //计算天数后剩余的毫秒数
	hours := math.Floor(float64(leave1 / 3600))
	//beego.Info(days, hours)

	// //计算相差分钟数
	leave2 := uint64(leave1) % 3600 //计算小时数后剩余的毫秒数
	minutes := math.Floor(float64(leave2 / 60))
	beego.Info(days, hours, minutes)

	// //计算相差秒数
	leave3 := float64(leave2 % 60) //计算分钟数后剩余的毫秒数
	seconds := math.Round(leave3)
	beego.Info(days, hours, minutes, seconds)

	//unix时间戳转换为UTC时间戳
	timeutc := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
	//获取时间差
	timesub := ""
	if days > 0 {
		timesub = strconv.Itoa(int(days)) + "天前"
	} else if hours > 0 {
		timesub = strconv.Itoa(int(hours)) + "小时前"
	} else if minutes > 0 {
		timesub = strconv.Itoa(int(minutes)) + "分钟前"
	} else if seconds > 0 {
		timesub = strconv.Itoa(int(seconds)) + "秒前"
	}
	return timeutc, timesub
}

//显示文章详情
//传入文章id,显示文章详细内容
func (c *WriteToolController) ShowArticleDetail() {
	articleInfo := ArticleInfo{}
	article_id := c.GetString(":id")
	if c.isLogin {
		beego.Info("一登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
	}

	//获取文章详细内容
	article_detail := models.Article_bases{}
	utils.DB.Where("id=?", article_id).Find(&article_detail)
	//fmt.Println(all_daliy_logs[3].Content)
	article_class_id := article_detail.ClassId
	articleInfo.Id, _ = strconv.Atoi(article_id)

	//获取文章类别
	article_class := models.Article_classes{}
	utils.DB.Where("id=?", article_class_id).Find(&article_class)

	articleInfo.ArticleFirstClass = article_class.FirstClass
	articleInfo.ArticleSecondClass = article_class.SecondClass
	articleInfo.ArticleContent = article_detail.Content
	articleInfo.ArticleAuthor = article_detail.Author
	articleInfo.ArticleTitle = article_detail.Title

	//获取文章类别路由
	articleInfo.ArticleFirstClassRoute = "/" + article_class.FirstClassRoute
	articleInfo.ArticleSecondClassRoute = "/" + article_class.FirstClassRoute + "/" + article_class.SecondClassRoute

	//获取文章发布及更新时间段
	///////////////////////////////
	articleInfo.ArticleCreateTimeUTC, articleInfo.ArticleCreateTimeSub = TimeTrans(article_detail.CreatedAt)
	articleInfo.ArticleUpdateTimeUTC, articleInfo.ArticleUpdateTimeSub = TimeTrans(article_detail.UpdatedAt)

	c.Data["articleInfo"] = articleInfo
	//beego.Info(article_detail.Content)

	c.Layout = "layouts/app.tpl"
	c.TplName = "pages/article_detail.tpl"
}

type ArticleId struct {
	Id int64 `json:id`
}

type DataReback struct {
	Mes string
}

//删除文章
func (c *WriteToolController) DeleteArticle() {
	articleid := string(c.Ctx.Input.RequestBody)
	beego.Info(articleid)
	articleids := strings.Split(articleid, "=")
	beego.Info(articleids[1])

	//删除文章操作需要用户为已登陆用户，如果没有登陆，则跳转

	if c.isLogin {
		beego.Info("一登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		// 判断用户是否是管理员账户
		// 1. 是管理员账户
		utils.DB.Where("id = ?", articleids[1]).Delete(&models.Article_bases{})
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
	}
	c.ServeJSON()
}
