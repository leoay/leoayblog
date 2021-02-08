package utils

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpasswd")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqlserver := beego.AppConfig.String("mysqlserver")
	mysqldb := beego.AppConfig.String("mysqldb")

	DB, err = gorm.Open("mysql", mysqluser+":"+mysqlpass+"@("+mysqlserver+":"+mysqlport+")/"+mysqldb+"?charset=utf8")
	if err != nil {
		beego.Error(err)
		return
	}
}
