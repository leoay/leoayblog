package controllers

import (
	"WriterReaderLog/models"
	"WriterReaderLog/utils"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type CreatorController struct {
	UserStatusController
}

type ArticleMessage struct {
	Id          uint
	Title       string
	Author      string
	EditorId    uint
	ClassId     uint
	Content     string
	ApprovalNum int
	StarNum     int
	IsTemp      bool
	MainPicPath string
	Summary     string
	Content_md  string
	Comments    string
	CreatedAt   uint64
	UpdatedAt   uint64
	TimeUTC     string
	TimeSub     string
}

//公众号文章
//首页
func (c *CreatorController) WeArticleIndex() {
	article_message := ArticleMessage{}
	article_message_arr := []ArticleMessage{}
	article_lists := []models.Article_bases{}
	utils.DB.Where("class_id=?", 7).Order("id desc").Find(&article_lists)
	//utils.DB.Where("class_id=?", 13).Order("id desc").Find(&article_lists)
	//后端分页功能实现
	var count int
	count = int(len(article_lists))

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

	article_lists = []models.Article_bases{}
	utils.DB.Where("class_id=?", 7).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&article_lists)

	c.Data["article_lists"] = article_lists
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
	var pagetop_arr []int
	var pagelast_arr []int

	if pagecount > 7 {
		for i := 0; i < 4; i++ {
			pagetop_arr = append(pagetop_arr, i+1)
		}
		for i := pagecount - 3; i < pagecount; i++ {
			pagelast_arr = append(pagelast_arr, i+1)
		}
		c.Data["pagelast_arr"] = pagelast_arr
	} else {
		for i := 0; i < pagecount; i++ {
			pagetop_arr = append(pagetop_arr, i+1)
		}
	}
	c.Data["pagetop_arr"] = pagetop_arr
	beego.Info(pagetop_arr)

	for _, artMes := range article_lists {
		//beego.Info(index)
		article_message.Id = artMes.Id
		article_message.Title = artMes.Title
		article_message.Author = artMes.Author
		article_message.EditorId = artMes.EditorId
		article_message.ClassId = artMes.ClassId
		article_message.Content = artMes.Content
		article_message.ApprovalNum = artMes.ApprovalNum
		article_message.StarNum = artMes.StarNum
		article_message.IsTemp = artMes.IsTemp
		article_message.Content_md = artMes.Content_md
		article_message.Comments = artMes.Comments
		article_message.MainPicPath = artMes.MainPicPath
		article_message.Summary = artMes.Summary
		article_message.CreatedAt = artMes.CreatedAt
		article_message.UpdatedAt = artMes.UpdatedAt

		///////////////////////////////
		timeSubUnix := uint64(time.Now().Unix()) - artMes.CreatedAt
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
		article_message.TimeUTC = time.Unix(int64(artMes.CreatedAt), 0).Format("2006-01-02 15:04:05")
		//获取时间差
		if days > 0 {
			article_message.TimeSub = strconv.Itoa(int(days)) + "天前"
		} else if hours > 0 {
			article_message.TimeSub = strconv.Itoa(int(hours)) + "小时前"
		} else if minutes > 0 {
			article_message.TimeSub = strconv.Itoa(int(minutes)) + "分钟前"
		} else if seconds > 0 {
			article_message.TimeSub = strconv.Itoa(int(seconds)) + "秒前"
		}
		article_message_arr = append(article_message_arr, article_message)
	}

	c.Data["article_lists"] = article_message_arr
	//beego.Info(article_lists[5])

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
	c.TplName = "pages/create/we_article.tpl"
}

//流水账，碎碎念
func (c *CreatorController) DaliyThinkIndex() {
	article_lists := []models.Article_bases{}
	utils.DB.Where("class_id=?", 13).Order("id desc").Find(&article_lists)

	//后端分页功能实现
	var count int
	count = int(len(article_lists))
	if count == 0 {
		count = 1
	}

	pageindex, _ := strconv.Atoi(c.GetString("pageIndex"))
	pagesize := 3
	pagecount := count / pagesize
	if pagecount == 0 {
		pagecount = 1
	}
	article_lists = []models.Article_bases{}
	utils.DB.Where("class_id=?", 13).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&article_lists)

	c.Data["article_lists"] = article_lists
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
	var pagetop_arr []int
	var pagelast_arr []int

	if pagecount > 7 {
		for i := 0; i < 4; i++ {
			pagetop_arr = append(pagetop_arr, i+1)
		}
		for i := pagecount - 3; i < pagecount; i++ {
			pagelast_arr = append(pagelast_arr, i+1)
		}
		c.Data["pagelast_arr"] = pagelast_arr
	} else {
		for i := 0; i < pagecount; i++ {
			pagetop_arr = append(pagetop_arr, i+1)
		}
	}
	c.Data["pagetop_arr"] = pagetop_arr

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
	c.TplName = "pages/create/daliy_note.tpl"
}
