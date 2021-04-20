package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpasswd")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqlserver := beego.AppConfig.String("mysqlserver")
	mysqldb := beego.AppConfig.String("mysqldb")

	dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		beego.Error(err)
		return
	}
	return
}

func closeDB() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()
}
