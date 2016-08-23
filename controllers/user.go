package controllers

import (
	"time"

	m "Personal-Dictionary/models"

	"github.com/astaxie/beego"
)

// UserController Login & Logout & SignUp.
type UserController struct {
	beego.Controller
}

// Login do Login or Show Login Page.
func (c *UserController) Login() {
	isAjax := c.GetString("isajax")

	if isAjax == "1" {
		username := c.GetString("username")
		password := c.GetString("password")

		user, err := doLogin(username, password)
		if err == nil {
			c.SetSession("userinfo", user)
			c.Resp(true, "登陆成功")
			return
		}
		c.Resp(false, err.Error())
		return
	}
	// Login Fail! relogin.
	c.Data["Title"] = beego.AppConfig.String("login_title")
	c.TplName = "login.tpl"
}

func doLogin(username, password string) (m.User, error) {
	return m.User{
		Id:         1,
		Username:   "HackerZ",
		Password:   "hackerzgz",
		Salt:       "zgz",
		Createtime: time.Now(),
	}, nil
}

// func SignUp() {

// }

// Logout ...
func (c *UserController) Logout() {
	c.DelSession("userinfo")
	c.Ctx.Redirect(302, "/login")
}

// Resp Response request Status and Info.
func (c *UserController) Resp(status bool, str string) {
	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	c.ServeJSON()
}
