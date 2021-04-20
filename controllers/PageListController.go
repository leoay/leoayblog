package controllers

//文章列表控制器, 不论什么类型的文章都使用同一个文章列表模板

import (
	"WriterReaderLog/models"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//PageListController 文章列表页
type PageListController struct {
	AuthController
}

//ArticleMessage 发送到前端的文章结构体信息
type ArticleMessage struct {
	ID          int64
	Title       string
	Author      string
	ClassID     int64
	Content     string
	ApprovalNum int
	StarNum     int
	IsTemp      bool
	MainPicPath string
	Summary     string
	ContentMd   string
	Comments    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	TimeUTC     time.Time
	TimeSub     time.Time
}

//GetPageList 获取文章列表，根据不同的类别筛选不同的文章
func (c *PageListController) GetPageList() {

	articleMessage := ArticleMessage{}
	articleMessageArr := []ArticleMessage{}
	articleLists := []models.Articles{}

	var classID = c.GetString("classId")

	//先获取文章类别
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
	db.Where("class_id=?", classID).Order("id desc").Find(&articleLists)

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

	db.Order("id desc").Where("class_id=?", 1).Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&articleLists)
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close() //上传成功后显示信息

	c.Data["article_lists"] = articleLists
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
		c.Data["pagelast_arr"] = pagelastArr
	} else {
		for i := 0; i < pagecount; i++ {
			pagetopArr = append(pagetopArr, i+1)
		}
	}
	c.Data["pagetop_arr"] = pagetopArr

	for _, artMes := range articleLists {

		articleMessage.ID = artMes.ID
		articleMessage.Title = artMes.Title
		articleMessage.Author = artMes.Author
		articleMessage.ClassID = artMes.ClassID
		articleMessage.Content = artMes.Content
		articleMessage.ApprovalNum = artMes.ApprovalNum
		articleMessage.StarNum = artMes.StarNum
		articleMessage.IsTemp = artMes.IsTemp

		articleMessage.MainPicPath = artMes.MainPicPATH
		articleMessage.Summary = artMes.Summary
		articleMessage.ContentMd = artMes.ContentMd
		articleMessage.Comments = artMes.Comments
		articleMessage.CreatedAt = artMes.CreatedAt
		articleMessage.UpdatedAt = artMes.UpdatedAt

		///////////////////////////////
		//timeSubUnix := time.Now().Unix() - artMes.CreatedAt
		timeSubUnix := 200
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
		// 	articleMessage.TimeSub = strconv.Itoa(int(days)) + "天前"
		// } else if hours > 0 {
		// 	articleMessage.TimeSub = strconv.Itoa(int(hours)) + "小时前"
		// } else if minutes > 0 {
		// 	articleMessage.TimeSub = (int(minutes)) + "分钟前"
		// } else if seconds > 0 {
		// 	article_message.TimeSub = strconv.Itoa(int(seconds)) + "秒前"
		// }
		articleMessageArr = append(articleMessageArr, articleMessage)
	}

	//前后端分离

	//c.Data["article_lists"] = article_message_arr

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

	c.Layout = "layouts/app.tpl"
	c.TplName = "pages/codeTec/tec_column.tpl"
}
