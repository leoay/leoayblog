package controllers

import (
	"WriterReaderLog/models"
	"WriterReaderLog/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type OpenPicPoolController struct {
	beego.Controller
}

// type ReturnMsg struct {
// 	Success int    `json:"success"`
// 	Message string `json:"message"`
// 	Url     string `json:"url"`
// }

func (c *OpenPicPoolController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "openpicpool.tpl"
}

//图片上传路由
func (c *OpenPicPoolController) OpenPicUpload() {
	//TODO: 加入鉴权
	//beego.Info(c.GetFile("file"))
	imgfile, header, err := c.GetFile("file")
	if err != nil {
		fmt.Println("getfile err ", err)
	}
	defer imgfile.Close()
	timeNow := time.Now().UnixNano() / 1e6
	filename := strings.TrimSpace(header.Filename)
	//beego.Info(filename)
	picture := strings.Split(filename, ".")
	layout := strings.ToLower(picture[len(picture)-1])
	imgpathNew := "static/picture_pool/" + strconv.FormatInt(timeNow, 10) + "." + layout
	fmt.Println(imgpathNew)
	err = c.SaveToFile("file", imgpathNew)
	if err != nil {
		fmt.Println("File upload failed！", err)
	} else {
		fmt.Println("File upload succeed!")
		//将图片信息保存到数据库                                                                                                                                                                                                                                                                                                                                      //上传成功后显示信息
		utils.DB.Create(&models.Image_Base{Name: filename, Path: imgpathNew, UserId: -1, CreatedAt: uint64(time.Now().Unix()), UpdatedAt: uint64(time.Now().Unix())}) // 通过数据的指针来创建
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

//图片集获取路由
func (c *OpenPicPoolController) OpenQueryImages() {
	beego.Info("sdsdsdsd")

	var image_lists = []models.Image_Base{}
	utils.DB.Where("user_id = ?", -1).Find(&image_lists)
	beego.Info(image_lists)
	c.Ctx.WriteString("sdsd")
	// 	if c.isLogin {
	// 		beego.Info("==================================", c.User.Id)
	// 		var image_lists = []models.Image_Base{}
	// 		utils.DB.Where("user_id = ?", c.User.Id).Find(&image_lists)
	// 		beego.Info(image_lists)

	// 		//后端分页功能实现
	// 		var count int
	// 		//如果只是获取数据长度的话，此处可以进一步优化
	// 		count = int(len(image_lists))
	// 		if count == 0 {
	// 			count = 1
	// 		}
	// 		beego.Info("count===", count)

	// 		pageindex, _ := strconv.Atoi(c.GetString("pageIndex"))
	// 		pagesize := 9
	// 		pagecount := count / pagesize
	// 		if pagecount == 0 {
	// 			pagecount = 1
	// 		}

	// 		imagesec_lists := []models.Image_Base{}
	// 		utils.DB.Where("user_id=?", c.User.Id).Order("id desc").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&imagesec_lists)

	// 		c.Data["pagecount"] = pagecount
	// 		c.Data["pageindex"] = pageindex
	// 		if pageindex > 1 {
	// 			c.Data["prepage"] = pageindex - 1
	// 		} else {
	// 			c.Data["prepage"] = pageindex
	// 		}

	// 		if pageindex < pagecount {
	// 			c.Data["nexpage"] = pageindex + 1
	// 		} else {
	// 			c.Data["nexpage"] = pageindex
	// 		}

	// 		//判断pagecount是否大于7
	// 		var pagetop_arr []int
	// 		var pagelast_arr []int
	// 		//九宫图图片集
	// 		if pagecount > 7 {
	// 			for i := 0; i < 4; i++ {
	// 				pagetop_arr = append(pagetop_arr, i+1)
	// 			}
	// 			for i := pagecount - 3; i < pagecount; i++ {
	// 				pagelast_arr = append(pagelast_arr, i+1)
	// 			}
	// 			c.Data["pagelast_arr"] = pagelast_arr
	// 		} else {
	// 			for i := 0; i < pagecount; i++ {
	// 				pagetop_arr = append(pagetop_arr, i+1)
	// 			}
	// 		}
	// 		c.Data["pagetop_arr"] = pagetop_arr

	// 		var imageMessRow ImageRows
	// 		var imageMessCol ImagesCols
	// 		var imageMessItem ImagesMessage

	// 		var imageMessItemArray []ImagesMessage
	// 		var imageMessColArray []ImagesCols

	// 		i := 0
	// 		j := 0
	// 		k := 0

	// 		for i = 0; i < int(count/3); i++ {
	// 			for j = 0; j < 3; j++ {
	// 				k++
	// 				imageMessItem.Name = imagesec_lists[k].Name
	// 				imageMessItem.Url = imagesec_lists[k].Path
	// 				imageMessItemArray = append(imageMessItemArray, imageMessItem)
	// 			}
	// 			imageMessCol.ImagesCol = imageMessItemArray
	// 			imageMessColArray = append(imageMessColArray, imageMessCol)
	// 			imageMessItemArray = []ImagesMessage{}
	// 		}
	// 		imageMessRow.ImageRow = imageMessColArray
	// 		imageMessItemArray = []ImagesMessage{}
	// 		//imageMessColArray = []ImagesCols{}

	// 		for i = (int(count/3) * 3); i < count; i++ {
	// 			imageMessItem.Name = imagesec_lists[i].Name
	// 			imageMessItem.Url = imagesec_lists[i].Path
	// 			imageMessItemArray = append(imageMessItemArray, imageMessItem)
	// 		}
	// 		imageMessCol.ImagesCol = imageMessItemArray
	// 		imageMessColArray = append(imageMessColArray, imageMessCol)
	// 		imageMessRow.ImageRow = imageMessColArray

	// 		jsonBytes, err := json.Marshal(imageMessRow)
	// 		if err != nil {
	// 			beego.Info(err)
	// 		}
	// 		beego.Info(string(jsonBytes))
	// 		c.Data["json"] = string(jsonBytes)
	// 		c.ServeJSON()
	// 	} else {
	// 		c.Redirect("/", 302)
	// 	}
}
