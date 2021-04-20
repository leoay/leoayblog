package main

import (
	"WriterReaderLog/models"
	_ "WriterReaderLog/models"
	_ "WriterReaderLog/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalSessions *session.Manager

func init() {
	//utils.DB.Exec("drop table article_read_record")

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
	// db.Where("name=?", username).Find(&user)
	beego.Info("Start Drop table Users")
	db.Exec("drop table users")
	db.AutoMigrate(&models.Articles{})
	db.AutoMigrate(&models.Images{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ArticleClass{})

	//db.Delete(models.Article_classes{}, "first_class LIKE ?", "%*%")
	//删除全部
	//utils.DB.Where("id LIKE ?", "%%").Delete(models.Article_classes{})
	//utils.DB.Exec("truncate table article_bases")
	//db.Exec("truncate table article_classes")

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}

	//utils.DB.Exec("truncate table users")

	// 读取目录配置json文件
	// jsonFile, err := os.Open("static/topic_class_conf.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	//解析json
	// type NameRoute struct {
	// 	ClassName string `json:"class_name"`
	// 	RouteName string `json:"route_name"`
	// }

	// type Mon_node struct {
	// 	Name      NameRoute   `json:"name"`
	// 	Son_class []NameRoute `json:"son_class"`
	// }

	// type Mon_class struct {
	// 	Mon_class []Mon_node `json:mon_class`
	// }

	// var mon_class Mon_class
	// json.Unmarshal(byteValue, &mon_class)

	// for _, value := range mon_class.Mon_class {
	// 	for _, value2 := range value.Son_class {
	// 		article_class := models.Article_classes{FirstClass: value.Name.ClassName, SecondClass: value2.ClassName, FirstClassRoute: value.Name.RouteName, SecondClassRoute: value2.RouteName}
	// 		utils.DB.Create(&article_class)
	// 	}
	// }

	//article_class := models.Article_classes{FirstClass: "代码技术", SecondClass: "技术专栏"}
	db.Exec("truncate table users")

	pwd, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
	userI := models.User{Name: "leoay", Password: string(pwd)}
	db.Create(&userI) // 通过数据的指针来创建
	sqlDB.Close()
	// pwd, _ = bcrypt.GenerateFromPassword([]byte("panxueshen"), bcrypt.DefaultCost)
	// userI = models.User{Name: "panxueshen", Pwd: string(pwd)}
	// utils.DB.Create(&userI) // 通过数据的指针来创建

	//db.Delete(models.Create_daliy_thinks{}, "id LIKE ?", "%*%")
	//Session初始化
	// sessionConfig := &session.ManagerConfig{
	// 	CookieName:      "gosessionid",
	// 	EnableSetCookie: true,
	// 	Maxlifetime:     3600,
	// 	Secure:          true,
	// 	CookieLifeTime:  3600,
	// 	ProviderConfig:  "./runtime/session",
	// }
	// globalSessions, _ = session.NewManager("file", sessionConfig)
	// go globalSessions.GC()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
}

func main() {
	beego.Run()
}
