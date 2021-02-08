package controllers

import (
	"WriterReaderLog/models"
	"WriterReaderLog/utils"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

//UserStatusController   用户状态结构体
type UserStatusController struct {
	beego.Controller
	isLogin bool        // 登陆状态标记
	User    models.User // 登陆的用户
}

//Prepare   用户登陆状态
func (c *UserStatusController) Prepare() {
	// 设置默认值
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

	user := models.User{Id: userid}
	utils.DB.Find(&user)
	// 判断密码是否修改过
	err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password))
	if err != nil {
		return
	}
	// 最后才设置登陆状态
	c.isLogin = true
	c.Data["isLogin"] = true
	c.User = user
}
