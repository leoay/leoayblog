package controllers

import (
	"WriterReaderLog/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//TODO 获取Session
func (c *SessionController) GetSessionData() {
	// beego.Info("connect success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	// user := models.User{}
	// //user.Name="浩秦"

	// resp["errno"] = 4001
	// resp["errmsg"] = "数据获取失败"
	// name := c.GetSession("name")
	// if name != nil {
	// 	user.Name = name.(string)
	// 	resp["errno"] = 0
	// 	resp["mrrmsg"] = "ok"
	// 	resp["data"] = user
	// }
	resp["data"] = "sdsd"
}

//TODO 删除对应的Session
func (c *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	c.DelSession("name")
	resp["errno"] = 0
	resp["errmsg"] = "ok"
}

// TODO 登录
func (c *SessionController) Login() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	//得到用户信息获取前端传递过来的json数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	//beego.Info(&resp)
	//判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil {
		fmt.Println("参数错误")
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}
	//与数据库匹配账号密码是否正确
	fmt.Println("与数据库匹配账号密码是否正确")
	o := orm.NewOrm()
	user := models.User{Name: resp["mobile"].(string)}
	moble := resp["mobile"].(string)
	qs := o.QueryTable("user")
	fmt.Println("----------------------")
	err := qs.Filter("mobile", moble).One(&user)
	fmt.Println("sdsdsdsdsdsdsd")
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		fmt.Println("用户或者密码错误")
		beego.Info("name=", resp["mobile"], "========password====", resp["password"])
		return
	}
	beego.Info(resp["password"])
	if user.Pwd != resp["password"] {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		fmt.Println("用户或者密码错误")
		beego.Info("3333name=", resp["mobile"], "========password====", resp["password"])
		return
	}
	//添加Session
	c.SetSession("name", resp["mobile"])
	c.SetSession("mobile", resp["mobile"])
	c.SetSession("user_id", user.Id)
	//返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.Ctx.WriteString("dsasdad")
}
func (c *SessionController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}
