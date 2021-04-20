package controllers

import (
	"WriterReaderLog/models"
	"fmt"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//LoginController 控制器结构
type LoginController struct {
	AuthController
}

//Index 登陆界面
func (c *LoginController) Index() {
	//Todo: 如果当前账户已登陆，则跳转到首页
	if c.isLogin {
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "block"
		c.Ctx.Redirect(302, "/")
	} else {
		c.Data["is_block"] = "none"
		c.Data["is_load_block"] = "none"
		c.Layout = "layouts/app.html"
		c.TplName = "auth/login.html"
	}
}

//Login 账户登陆
func (c *LoginController) Login() {
	beego.Info("SDADA")
	//判断session是否有效
	username := c.GetString("username")
	password := c.GetString("password")
	user := models.User{}

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
	db.Where("name=?", username).Find(&user)

	//查询完毕后，关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()
	beego.Info(user)
	//将密码与数据库中的密码进行对比
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		beego.Info("用户密码验证失败！")
	} else {
		beego.Info("用户密码验证通过！")
		c.SetSession("userid", user.ID)
		c.SetSession("password", password)

		//登陆成功，跳转到首页
		c.Ctx.Redirect(302, "/")
	}
}

// //tiken信息结构体
// type AuthRes struct {
// 	Access_token string `json:"access_token"`
// 	Scope        string `json:"scope"`
// 	Token_type   string `json:"token_type"`
// }

// //github用户信息结构体
// type UserInfo struct {
// 	Login     string `json:"login"`
// 	Id        string `json:"id"`
// 	NodeId    string `json:"node_id"`
// 	AvatarUrl string `json:"avatar_url"`
// }

// //普通用户登陆

// //Github 登陆，需要判断当前系统账户中是否有该github账号对应的用户信息
// func (c *LoginController) Github_oauth() {
// 	sessionInfo := c.GetSession("userId")
// 	if sessionInfo != nil {
// 		c.Ctx.Redirect(302, "/")
// 	} else {
// 		code := c.GetString("code")
// 		fmt.Println(code)
// 		//Todo: 如果当前账户已登陆，则跳转到首页

// 		url := "https://github.com/login/oauth/access_token?client_id=d0c10af3bba71a8d6753&client_secret=93ae985f1d99cf4e37b1276dfa8db35a1015aa40&code=" + code
// 		req, _ := http.NewRequest("GET", url, nil)
// 		req.Header.Set("Accept", "application/json")
// 		resp, err := (&http.Client{}).Do(req)
// 		if err != nil {

// 		}
// 		defer resp.Body.Close()
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		fmt.Println(string(body))

// 		authRes := AuthRes{}
// 		err = json.Unmarshal(body, &authRes)
// 		fmt.Println(authRes)
// 		access_token := authRes.Access_token
// 		req, _ = http.NewRequest("GET", "https://api.github.com/user", nil)
// 		req.Header.Set("Authorization", "token "+access_token)
// 		req.Header.Set("Accept", "application/json")
// 		resp, err = (&http.Client{}).Do(req)
// 		if err != nil {

// 		}
// 		defer resp.Body.Close()
// 		respByte, _ := ioutil.ReadAll(resp.Body)
// 		//resp_str := string(respByte)
// 		userinfo := UserInfo{}
// 		err = json.Unmarshal(respByte, &userinfo)
// 		fmt.Println(userinfo)
// 		name := userinfo.Login
// 		id := userinfo.Id
// 		avata := userinfo.AvatarUrl
// 		c.Ctx.SetCookie("status", "success", 1)
// 		c.Ctx.SetCookie("message", "登陆成功，开始愉快浏览把！", 1)
// 		c.SetSession("userId", base64.StdEncoding.EncodeToString([]byte(name+"##"+id+"##"+avata)))
// 	}
// 	c.Ctx.Redirect(302, "/")
// }

//微信登陆，直接微信扫码登陆

// func (c *UserLoginController) Post() {
// 	// 创建logs（记录器）
// 	l := logs.BeeLogger{}
// 	l.SetPrefix("用户登陆API：")
// 	// 判断是否已经登陆 - 已登陆就终结登陆流程
// 	if c.isLogin == true {
// 		l.Debug("用户已登陆")
// 		c.Ctx.ResponseWriter.WriteHeader(202)
// 		// util.CreateMsg - 这个是我自己写的方法，请参考我博文下面的源代码
// 		c.Data["json"] = util.CreateMsg(202, "用户已登陆")
// 		c.ServeJSON()
// 		return
// 	}
// 	// 用户登陆
// 	username := strings.TrimSpace(c.GetString("username"))
// 	password := strings.TrimSpace(c.GetString("password"))
// 	// 验证表单数据 - 这里使用官方表单数据验证包（结合自定义验证）
// 	// github.com/astaxie/beego/validation
// 	// 可以下面附加的验证工具，或者直接把这一段删掉
// 	u := &valid.ValidateUser{
// 		Username: username,
// 		Password: password,
// 	}
// 	err := u.ValidUser()
// 	if err != nil {
// 		l.Debug("用户注册失败，原因:%s", err.Error())
// 		c.Ctx.ResponseWriter.WriteHeader(403)
// 		c.Data["json"] = util.CreateMsg(403, err.Error())
// 		c.ServeJSON()
// 		return
// 	}
// 	// 创建User对象 - 不设置密码
// 	user := models.User{Username: username}
// 	// 创建ORM对象
// 	o := orm.NewOrm()
// 	// 读取用户对象 - 这里使用One方法而不使用Read可以避免读取整个User对象，稍微优化
// 	// 虽然在这里没有什么卵用，但是如果用户模型比较复杂的情况下还是有效果的
// 	err = o.QueryTable("user").Filter("username", &user.Username).One(&user, "id","password")
// 	if err != nil {
// 		l.Debug("用户不存在")
// 		c.Ctx.ResponseWriter.WriteHeader(401)
// 		c.Data["json"] = util.CreateMsg(401, "用户不存在")
// 		c.ServeJSON()
// 		return
// 	}
// 	// 验证用户密码
// 	// 第一个参数是数据库存储的加密密码，第二个参数是原密码
// 	// 密码是单方面加密，不可解密
// 	// bcrypt包提供的安全加密可以应对大部分的安全要求，具体实现细节请自行百度
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		l.Debug("用户密码错误")
// 		c.Ctx.ResponseWriter.WriteHeader(401)
// 		c.Data["json"] = util.CreateMsg(401, "密码错误")
// 		c.ServeJSON()
// 		return
// 	}
// 	// 设置服务器Session存储
// 	c.SetSession("userid", user.Id)
// 	c.SetSession("password", password) // 这里存储的是原始密码

// 	// 返回成功信息
// 	c.Data["json"] = util.CreateMsg(200, "用户登陆成功")
// 	c.ServeJSON()
// }
