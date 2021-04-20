package controllers

//写作工具控制器

import (
	"WriterReaderLog/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//WriteToolController 控制器结构体
type WriteToolController struct {
	AuthController
}

//Index 写作工具首页
func (c *WriteToolController) Index() {
	if c.isLogin {
		beego.Info("已登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		c.Data["newarticle"] = 1
		c.Layout = "layouts/app.html"
		c.TplName = "pages/writeTool/all_writer.html"
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	}
}

//SaveTemp 保存草稿入口
func (c *WriteToolController) SaveTemp() {
	articleID := c.GetString("articleid")
	var articleClass models.ArticleClass
	markdownText := c.GetString("leoaylog-editor-markdown-doc")
	title := c.GetString("title")
	firstClass := c.GetString("article_class1")
	secondClass := c.GetString("article_class2")
	htmlCode := c.GetString("leoaylog-editor-html-code")
	summary := c.GetString("summary")
	mainPicPath := c.GetString("main_pic_path")

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpasswd")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqlserver := beego.AppConfig.String("mysqlserver")
	mysqldb := beego.AppConfig.String("mysqldb")

	dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		beego.Error(err)
		return
	}

	db.Where("first_class = ? AND second_class = ?", firstClass, secondClass).Find(&articleClass)

	if articleID != "" {
		//编辑原来的草稿文章

		articleBase := models.Articles{Title: title, ClassID: articleClass.ID, Content: htmlCode, ContentMd: markdownText, Summary: summary, MainPicPATH: mainPicPath, IsTemp: true}
		db.Model(&models.Articles{}).Where("id = ?", articleID).Updates(articleBase)
	} else {
		//创建的草稿文章
		articleBase := models.Articles{Title: title, ClassID: articleClass.ID, Content: htmlCode, ContentMd: markdownText, Summary: summary, MainPicPATH: mainPicPath, IsTemp: true}
		db.Create(&articleBase)
	}

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()

	//保存草稿后，返回到编辑页面
	c.Ctx.SetCookie("status", "success", 1)
	c.Ctx.SetCookie("message", "恭喜！保存草稿成功！", 1)
	c.Ctx.Redirect(302, "/")
}

//Publish 发布文章
func (c *WriteToolController) Publish() {
	//登录之后才进行保存
	if c.isLogin {
		//fmt.Println("-------------------", c.GetString("upload-url"))
		newarticle := c.GetString("newarticle")
		//获取文章编辑表单
		markdownText := c.GetString("leoaylog-editor-markdown-doc")
		title := c.GetString("title")
		firstClass := c.GetString("article_class1")
		secondClass := c.GetString("article_class2")
		summary := c.GetString("summary")
		mainPicPath := c.GetString("upload-url")

		htmlCode := c.GetString("leoaylog-editor-html-code")
		var articleClass models.ArticleClass

		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpasswd")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqlserver := beego.AppConfig.String("mysqlserver")
		mysqldb := beego.AppConfig.String("mysqldb")

		dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			beego.Error(err)
			return
		}

		//查询完毕后，关闭连接
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		defer sqlDB.Close()

		db.Where("first_class = ? AND second_class = ?", firstClass, secondClass).Find(&articleClass)
		articleID := c.GetString("articleid")
		if articleID != "" {
			//编辑原来的草稿文章
			articleBase := models.Articles{Title: title, ClassID: articleClass.ID, Content: htmlCode, ContentMd: markdownText, Summary: summary, MainPicPATH: mainPicPath, IsTemp: true}
			db.Save(&articleBase)
		} else {
			//创建的草稿文章
			articleBase := models.Articles{Title: title, ClassID: articleClass.ID, Content: htmlCode, ContentMd: markdownText, Summary: summary, MainPicPATH: mainPicPath, IsTemp: false}
			db.Create(&articleBase)
		}

		istemp := false
		if newarticle == "1" {
			db.Create(&models.Articles{Title: title, Author: c.User.Name, MainPicPATH: mainPicPath, Summary: summary, UserID: c.User.ID, ClassID: articleClass.ID, Content: htmlCode, ApprovalNum: 10, StarNum: 10, IsTemp: istemp, ContentMd: markdownText, Comments: "leoay"}) // 通过数据的指针来创建
		} else {
			ariid, _ := strconv.Atoi(c.GetString("articleid"))
			db.Model(&models.Articles{ID: int64(ariid)}).Updates(models.Articles{Title: title, ClassID: articleClass.ID, MainPicPATH: mainPicPath, Summary: summary, Content: htmlCode, ApprovalNum: 10, StarNum: 10, IsTemp: istemp, ContentMd: markdownText, Comments: "leoay"})
		}
	}
	c.Ctx.SetCookie("status", "success", 1)
	c.Ctx.SetCookie("message", "恭喜！你已成功发布一篇文章，快写下一篇吧！", 1)
	c.Ctx.Redirect(302, "/")
}

//EditArticle 重新编辑文章
func (c *WriteToolController) EditArticle() {
	artID := c.GetString(":id")
	if c.isLogin {
		var articles models.Articles

		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpasswd")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqlserver := beego.AppConfig.String("mysqlserver")
		mysqldb := beego.AppConfig.String("mysqldb")

		dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			beego.Error(err)
			return
		}

		//查询用户是否存在
		db.Where("id=?", artID).Find(&articles)

		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		c.Data["markdown"] = articles.ContentMd
		c.Data["artTitle"] = articles.Title
		c.Data["articleid"] = articles.ID
		c.Data["newarticle"] = 0
		c.Data["main_pic_path"] = articles.MainPicPATH
		c.Data["summary"] = articles.Summary
		c.Data["istemp"] = articles.IsTemp
		classid := articles.ClassID
		var classT models.ArticleClass
		db.Where("id=?", classid).Find(&classT)

		//查询完毕后，关闭连接
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		sqlDB.Close()

		c.Data["firstClass"] = classT.FirstClass
		c.Data["secondClass"] = classT.SecondClass

		c.Layout = "layouts/app.html"
		c.TplName = "pages/writeTool/all_writer.html"
	} else {
		//如果session超时，则显示登陆按钮
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	}
}

//ArticleInfo 文章信息
type ArticleInfo struct {
	ID                      int64
	ArticleTitle            string
	ArticleContent          string
	ArticleFirstClass       string
	ArticleSecondClass      string
	ArticleAuthor           string
	ArticleReadNum          int
	ArticleCreateTimeSub    string
	ArticleUpdateTimeSub    string
	ArticleCreateTimeUTC    string
	ArticleUpdateTimeUTC    string
	ArticleFirstClassRoute  string
	ArticleSecondClassRoute string
}

//TimeTrans 计算时间工具
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

//ShowArticleDetail 显示文章详情
func (c *WriteToolController) ShowArticleDetail() {
	articleInfo := ArticleInfo{}
	articleID := c.GetString(":id")
	//user_ip := exnet.ClientPublicIP(c.Ctx.Request)

	beego.Info("11111111111111111111111111", c.AppController)

	//addr := req.RemoteAddr // "IP:port" "192.168.1.150:8889"
	//判断当前作者是否已经读过这篇文章
	//article_record := models.Article_Read_Record{}
	//utils.DB.Where("article_id=?", article_id).Where("user_ip=?", user_ip).Find(&article_record)
	//beego.Info("@@@@@@@@@@@@@@@@@@@@", article_record)

	// if ("" == article_record.UserIp) && ("true" != article_id) {
	// 	//当前用户没有读过这篇文章，阅读数加1
	// 	utils.DB.Create(&models.Article_Read_Record{UserIp: user_ip, ArticleId: article_id}) // 通过数据的指针来创建
	// 	//更新阅读数
	// 	article_base := models.Article_bases{}

	// 	utils.DB.Where("id=?", article_id).Find(&article_base)
	// 	beego.Info("999999999999999999999999999999", article_base)
	// 	currentReadNum := article_base.ReadNum
	// 	beego.Info("555555555555555555555555555", currentReadNum)
	// 	currentReadNum = currentReadNum + 1
	// 	ariid, _ := strconv.Atoi(c.GetString("articleid"))
	// 	utils.DB.Model(&models.Article_bases{Id: uint(ariid)}).Updates(models.Article_bases{ReadNum: currentReadNum})
	// }

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
	articleDetail := models.Articles{}

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpasswd")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqlserver := beego.AppConfig.String("mysqlserver")
	mysqldb := beego.AppConfig.String("mysqldb")

	dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		beego.Error(err)
		return
	}

	//查询用户是否存在
	db.Where("id=?", articleID).Find(&articleDetail)

	articleClassID := articleDetail.ClassID
	articleInfo.ID = articleDetail.ClassID

	//获取文章类别
	articleClass := models.ArticleClass{}

	db.Where("id=?", articleClassID).Find(&articleClass)

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()

	articleInfo.ArticleFirstClass = articleClass.FirstClass
	articleInfo.ArticleSecondClass = articleClass.SecondClass
	articleInfo.ArticleContent = articleDetail.Content
	articleInfo.ArticleAuthor = articleDetail.Author
	articleInfo.ArticleTitle = articleDetail.Title

	articleInfo.ArticleReadNum = articleDetail.ReadNum

	//获取文章类别路由
	articleInfo.ArticleFirstClassRoute = "/" + articleClass.FirstClassRoute
	articleInfo.ArticleSecondClassRoute = "/" + articleClass.FirstClassRoute + "/" + articleClass.SecondClassRoute

	//获取文章发布及更新时间段
	///////////////////////////////
	// articleInfo.ArticleCreateTimeUTC, articleInfo.ArticleCreateTimeSub = TimeTrans(articleDetail.CreatedAt)
	// articleInfo.ArticleUpdateTimeUTC, articleInfo.ArticleUpdateTimeSub = TimeTrans(articleDetail.UpdatedAt)
	articleInfo.ArticleCreateTimeUTC, articleInfo.ArticleCreateTimeSub = "100", "100"
	articleInfo.ArticleUpdateTimeUTC, articleInfo.ArticleUpdateTimeSub = "100", "100"

	c.Data["articleInfo"] = articleInfo

	c.Layout = "layouts/app.html"
	c.TplName = "pages/article_detail.html"
}

//ArticleID 文章ID
type ArticleID struct {
	ID int64 `json:id`
}

//DataReback 数据反馈
type DataReback struct {
	Mes string
}

//DeleteArticle 删除文章
func (c *WriteToolController) DeleteArticle() {
	articleid := string(c.Ctx.Input.RequestBody)
	articleids := strings.Split(articleid, "=")

	if c.isLogin {
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		// 判断用户是否是管理员账户,只有管理员账户才能删除文章
		// 1. 是管理员账户

		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpasswd")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqlserver := beego.AppConfig.String("mysqlserver")
		mysqldb := beego.AppConfig.String("mysqldb")

		dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			beego.Error(err)
			return
		}

		db.Where("id = ?", articleids[1]).Delete(&models.Articles{})

		//查询完毕后，关闭连接
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		sqlDB.Close()

	} else {
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
	}
	c.ServeJSON()
}
