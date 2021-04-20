package controllers

import (
	"WriterReaderLog/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//OpenPicPoolController 开放图片上传控制器
type OpenPicPoolController struct {
	AuthController
}

// type ReturnMsg struct {
// 	Success int    `json:"success"`
// 	Message string `json:"message"`
// 	Url     string `json:"url"`
// }

//Index 公共图片池主页
func (c *OpenPicPoolController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "openpicpool.html"
}

//OpenPicUpload 图片上传路由
func (c *OpenPicPoolController) OpenPicUpload() {
	//TODO: 加入鉴权
	imgfile, header, err := c.GetFile("file")
	if err != nil {
		fmt.Println("getfile err ", err)
	}
	defer imgfile.Close()
	timeNow := time.Now().UnixNano() / 1e6
	filename := strings.TrimSpace(header.Filename)

	picture := strings.Split(filename, ".")
	layout := strings.ToLower(picture[len(picture)-1])
	imgpathNew := "static/picture_pool/" + strconv.FormatInt(timeNow, 10) + "." + layout
	fmt.Println(imgpathNew)
	err = c.SaveToFile("file", imgpathNew)
	if err != nil {
		fmt.Println("File upload failed！", err)
	} else {
		fmt.Println("File upload succeed!")
		//将图片信息保存到数据库
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

		db.Create(&models.Images{Name: filename, Path: imgpathNew, UserID: c.User.ID}) // 通过数据的指针来创建
		//查询完毕后，关闭连接
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		sqlDB.Close() //上传成功后显示信息
	}

	res := ReturnMsg{
		1, "success", beego.AppConfig.String("appurl") + imgpathNew,
	}

	c.Data["json"] = res
	fmt.Println(res)
	resStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(resStr))
	c.Ctx.SetCookie("status", "success", 1)
	c.Ctx.SetCookie("message", "恭喜！上传图片处成功", 1)
	c.Ctx.WriteString(string(resStr))
}

//OpenQueryImages 图片集获取路由
func (c *OpenPicPoolController) OpenQueryImages() {
	var imageLists = []models.Images{}
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
	db.Where("user_id = ?", -1).Find(&imageLists)
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close() //上传成功后显示信息

	db.Where("user_id = ?", c.User.ID).Find(&imageLists)
	beego.Info(imageLists)
	c.Ctx.WriteString("sdsd")
}
