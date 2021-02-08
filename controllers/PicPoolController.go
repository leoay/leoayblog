package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type PicPoolController struct {
	beego.Controller
}

type ReturnMsg struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

//图片上传路由
func (c *PicPoolController) Upload() {
	//	{
	//		success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//		message : "提示的信息，上传成功或上传失败及错误信息等。",
	//		url     : "图片地址"        // 上传成功时才返回
	//	}
	//c.Ctx.WriteString("dfdfd")
	imgfile, header, err := c.GetFile("editormd-image-file")
	if err != nil {
		fmt.Println("getfile err ", err)
	}
	defer imgfile.Close()
	timeNow := time.Now().UnixNano() / 1e6
	filename := header.Filename
	picture := strings.Split(filename, ".")
	layout := strings.ToLower(picture[len(picture)-1])
	imgpathNew := "static/picture_pool/" + strconv.FormatInt(timeNow, 10) + "." + layout
	fmt.Println(imgpathNew)
	err = c.SaveToFile("editormd-image-file", imgpathNew)
	if err != nil {
		fmt.Println("File upload failed！", err)
	} else {
		fmt.Println("File upload succeed!") //上传成功后显示信息
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
