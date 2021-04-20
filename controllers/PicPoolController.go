package controllers

//图片池控制器，用户用户图片上传
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

//PicPoolController 图片池控制器结构体
type PicPoolController struct {
	AuthController
}

//ReturnMsg 上传成功后的返回消息结构体
type ReturnMsg struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

//Upload 图片上传路由
func (c *PicPoolController) Upload() {
	//TODO: 加入鉴权
	if c.isLogin {
		imgfile, header, err := c.GetFile("editormd-image-file")
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
		err = c.SaveToFile("editormd-image-file", imgpathNew)
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

			db.Create(&models.Images{Name: filename, UserID: c.User.ID, Path: imgpathNew}) // 通过数据的指针来创建
			//查询完毕后，关闭连接
			sqlDB, err := db.DB()
			if err != nil {
				fmt.Println(err)
			}
			sqlDB.Close()
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
	} else {
		c.Redirect("/", 302)
	}
}

//ImageResultList  ddsad
type ImageResultList struct {
	ImgMsgs     []models.Images `json:"image_list"`
	PageCount   int             `josn:"page_count"`
	PageIndex   int             `json:""page_index`
	PrePage     int             `json:"prepage"`
	NextPage    int             `json:"nextpage"`
	PageLastArr []int           `json:"pagelast_arr"`
	PageTopArr  []int           `json:"pagetop_arr"`
}

//QueryImages 图片集获取路由
func (c *PicPoolController) QueryImages() {
	if c.isLogin {
		var imageResultList = ImageResultList{}
		var imageLists = []models.Images{}

		var count int
		var prepage int
		var nextpage int
		//如果只是获取数据长度的话，此处可以进一步优化
		count = int(len(imageLists))
		if count == 0 {
			count = 1
		}

		pageindex, _ := strconv.Atoi(c.GetString("pageIndex"))
		pagesize := 9
		pagecount := count / pagesize
		if pagecount == 0 {
			pagecount = 1
		}

		imagesecLists := []models.Images{}
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

		db.Where("user_id=?", c.User.ID).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&imagesecLists)
		//查询完毕后，关闭连接
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		sqlDB.Close()

		imageResultList.PageCount = pagecount
		imageResultList.PageIndex = pageindex

		if pageindex > 1 {
			prepage = pageindex - 1
		} else {
			prepage = pageindex
		}

		if pageindex < pagecount {
			nextpage = pageindex + 1
		} else {
			nextpage = pageindex
		}

		imageResultList.PrePage = prepage
		imageResultList.NextPage = nextpage

		//判断pagecount是否大于7
		var pagetopArr []int
		var pagelastArr []int
		//九宫图图片集
		if pagecount > 7 {
			for i := 0; i < 4; i++ {
				pagetopArr = append(pagetopArr, i+1)
			}
			for i := pagecount - 3; i < pagecount; i++ {
				pagelastArr = append(pagelastArr, i+1)
			}
		} else {
			for i := 0; i < pagecount; i++ {
				pagetopArr = append(pagetopArr, i+1)
			}
		}

		imageResultList.PageTopArr = pagetopArr
		imageResultList.PageLastArr = pagelastArr
		imageResultList.ImgMsgs = imagesecLists

		jsonBytes, err := json.Marshal(imageResultList)
		if err != nil {
			beego.Info(err)
		}

		c.Data["json"] = string(jsonBytes)
		c.ServeJSON()
	} else {
		c.Redirect("/", 302)
	}
}
