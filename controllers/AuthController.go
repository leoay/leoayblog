package controllers

import (
	"WriterReaderLog/models"
	"fmt"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//AuthController 登录状态
type AuthController struct {
	beego.Controller
	isLogin bool        // 登陆状态标记
	User    models.User // 登陆的用户
}

//Prepare  设置用户登陆状态
func (c *AuthController) Prepare() {
	c.isLogin = false
	c.Data["isLogin"] = false
	c.User = models.User{}
	// 获取Session信息 - 注意这里返回的是interface类型
	useridInterface := c.GetSession("userid")
	beego.Info(useridInterface)
	userid, ok := useridInterface.(int64)
	beego.Info(userid, ok)
	// 优化性能，如果Id不存在就不需要再获取Password
	if !ok {
		return
	}
	// 接下来是验证Password是否被更新 - 防止数据库密码更新后仍保持Session
	passwordInterface := c.GetSession("password")
	password, _ := passwordInterface.(string)

	user := models.User{ID: userid}

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

	db.Find(&user)

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()

	// 判断密码是否修改过
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return
	}
	// 最后才设置登陆状态
	c.isLogin = true
	c.Data["isLogin"] = true
	c.User = user
}
