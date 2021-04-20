package controllers

import (
	"WriterReaderLog/models"
	"fmt"
	"math"
	"strconv"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//MainController   控制器结构体(类似于构造函数)
type MainController struct {
	AuthController
}

//Get 的撒旦
func (c *MainController) Get() {
	beego.Info("是否登陆", c.isLogin)
	//如果当前账户已登陆

	articleMessage := ArticleMessage{}
	articleMessageArr := []ArticleMessage{}
	articleLists := []models.Articles{}
	//后端分页功能实现
	var count int
	count = int(len(articleLists))

	beego.Info(count)

	pageindex, _ := strconv.Atoi(c.GetString("pageIndex"))
	pagesize := 3
	if count == 0 {
		count = 1
	}
	pagecount := count / pagesize
	if pagecount == 0 {
		pagecount = 1
	}

	articleLists = []models.Articles{}

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

	db.Order("id desc").Where("is_temp=?", false).Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&articleLists)

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()
	c.Data["article_lists"] = articleLists
	beego.Info(articleLists)
	c.Data["pagecount"] = pagecount
	c.Data["pageindex"] = pageindex
	if pageindex > 1 {
		c.Data["prepage"] = pageindex - 1
	} else {
		c.Data["prepage"] = pageindex
	}

	if pageindex < pagecount {
		c.Data["nexpage"] = pageindex + 1
	} else {
		c.Data["nexpage"] = pageindex
	}

	//判断pagecount是否大于7
	var pagetopArr []int
	var pagelastArr []int

	if pagecount > 7 {
		for i := 0; i < 4; i++ {
			pagetopArr = append(pagetopArr, i+1)
		}
		for i := pagecount - 3; i < pagecount; i++ {
			pagelastArr = append(pagelastArr, i+1)
		}
		c.Data["pagelastArr"] = pagelastArr
	} else {
		for i := 0; i < pagecount; i++ {
			pagetopArr = append(pagetopArr, i+1)
		}
	}
	c.Data["pagetopArr"] = pagetopArr
	beego.Info(pagetopArr)

	for _, artMes := range articleLists {
		articleMessage.ID = artMes.ID
		articleMessage.Title = artMes.Title
		articleMessage.Author = artMes.Author
		articleMessage.ClassID = artMes.ClassID
		articleMessage.Content = artMes.Content
		articleMessage.ApprovalNum = artMes.ApprovalNum
		articleMessage.StarNum = artMes.StarNum
		articleMessage.IsTemp = artMes.IsTemp
		articleMessage.ContentMd = artMes.ContentMd
		articleMessage.Comments = artMes.Comments

		articleMessage.MainPicPath = artMes.MainPicPATH

		articleMessage.Summary = artMes.Summary
		articleMessage.CreatedAt = artMes.CreatedAt
		articleMessage.UpdatedAt = artMes.UpdatedAt

		///////////////////////////////
		//timeSubUnix := uint64(time.Now().Unix()) - artMes.CreatedAt
		timeSubUnix := 2000
		beego.Info(timeSubUnix)
		//years := math.Floor(float64(timeSubUnix / (24 * 3600)))
		//var days = Math.floor(dateDiff / (24 * 3600 * 1000));//计算出相差天数
		days := math.Floor(float64(timeSubUnix / (24 * 3600)))

		leave1 := math.Floor(float64(timeSubUnix % (24 * 3600))) //计算天数后剩余的毫秒数
		hours := math.Floor(float64(leave1 / 3600))
		beego.Info(days, hours)

		//计算相差分钟数
		leave2 := uint64(leave1) % 3600 //计算小时数后剩余的毫秒数
		minutes := math.Floor(float64(leave2 / 60))
		beego.Info(days, hours, minutes)

		//计算相差秒数
		leave3 := float64(leave2 % 60) //计算分钟数后剩余的毫秒数
		seconds := math.Round(leave3)
		beego.Info(days, hours, minutes, seconds)

		//unix时间戳转换为UTC时间戳
		articleMessage.TimeUTC = artMes.CreatedAt
		//获取时间差
		// if days > 0 {
		// 	article_message.TimeSub = strconv.Itoa(int(days)) + "天前"
		// } else if hours > 0 {
		// 	article_message.TimeSub = strconv.Itoa(int(hours)) + "小时前"
		// } else if minutes > 0 {
		// 	article_message.TimeSub = strconv.Itoa(int(minutes)) + "分钟前"
		// } else if seconds > 0 {
		// 	article_message.TimeSub = strconv.Itoa(int(seconds)) + "秒前"
		// }
		articleMessageArr = append(articleMessageArr, articleMessage)
	}

	c.Data["article_lists"] = articleMessageArr

	if c.isLogin {
		beego.Info("已登陆")
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
	} else {
		//如果session超时，则显示登陆按钮
		beego.Info("未登陆")
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
	}

	if c.isLogin {
		c.Data["user"] = c.User
		c.Data["is_block"] = "block"
		c.Data["is_load_block"] = "none"
		c.Layout = "layouts/app.html"
		c.TplName = "pages/root.html"
	} else {
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Layout = "layouts/app.html"
		c.TplName = "pages/root.html"
	}
}
