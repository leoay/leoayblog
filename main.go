package main

import (
	"WriterReaderLog/models"
	_ "WriterReaderLog/models"
	_ "WriterReaderLog/routers"
	"WriterReaderLog/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

//TODO:
//做一个类似minidoc的应用，通过命令行install， 安装并迁移数据库
//=========================================================

func init() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//beego.ErrorController(&controllers.ErrorController{})
	// var Creatin = models.Create_daliy_thinks{}
	// db.First(&Creatin)
	// fmt.Println(Creatin.Label)
	//defer utils.DB.Close()

	//db.Exec("drop table article_bases")

	utils.DB.AutoMigrate(&models.Article_bases{})
	utils.DB.AutoMigrate(&models.Image_flows{})
	utils.DB.AutoMigrate(&models.User{})
	utils.DB.AutoMigrate(&models.Article_classes{})
	utils.DB.AutoMigrate(&models.Article_bases{})
	utils.DB.AutoMigrate(&models.Image_Base{})

	//db.Delete(models.Article_classes{}, "first_class LIKE ?", "%*%")
	//删除全部
	//utils.DB.Where("id LIKE ?", "%%").Delete(models.Article_classes{})
	//utils.DB.Exec("truncate table article_bases")
	utils.DB.Exec("truncate table article_classes")
	//utils.DB.Exec("truncate table users")

	// 读取目录配置json文件
	jsonFile, err := os.Open("static/topic_class_conf.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//解析json
	type NameRoute struct {
		ClassName string `json:"class_name"`
		RouteName string `json:"route_name"`
	}

	type Mon_node struct {
		Name      NameRoute   `json:"name"`
		Son_class []NameRoute `json:"son_class"`
	}

	type Mon_class struct {
		Mon_class []Mon_node `json:mon_class`
	}

	var mon_class Mon_class
	json.Unmarshal(byteValue, &mon_class)

	for _, value := range mon_class.Mon_class {
		for _, value2 := range value.Son_class {
			article_class := models.Article_classes{FirstClass: value.Name.ClassName, SecondClass: value2.ClassName, FirstClassRoute: value.Name.RouteName, SecondClassRoute: value2.RouteName}
			utils.DB.Create(&article_class)
		}
	}

	//article_class := models.Article_classes{FirstClass: "代码技术", SecondClass: "技术专栏"}
	// pwd, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
	// userI := models.User{Name: "leoay", Pwd: string(pwd)}
	// utils.DB.Create(&userI) // 通过数据的指针来创建

	// pwd, _ = bcrypt.GenerateFromPassword([]byte("panxueshen"), bcrypt.DefaultCost)
	// userI = models.User{Name: "panxueshen", Pwd: string(pwd)}
	// utils.DB.Create(&userI) // 通过数据的指针来创建

	//db.Delete(models.Create_daliy_thinks{}, "id LIKE ?", "%*%")
	//Session初始化
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Maxlifetime:     3600,
		Secure:          true,
		CookieLifeTime:  3600,
		ProviderConfig:  "./runtime/session",
	}
	globalSessions, _ = session.NewManager("file", sessionConfig)
	go globalSessions.GC()
}

func main() {
	beego.Run()
}
