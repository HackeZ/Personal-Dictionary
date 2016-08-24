package controllers

import (
	"YuXuan-Admin/utils"
	"errors"
	"log"

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
	log.Println("isAjax -->", isAjax)

	if isAjax == "1" {
		username := c.GetString("username")
		password := c.GetString("password")

		log.Println("username -->", username)
		log.Println("password -->", password)

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
	user := m.GetUserByUsername(username)

	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != utils.PassEncode(password, user.Salt) {
		return m.User{}, errors.New("密码错误")
	}
	return user, nil
}

// SignUp 注册新用户
func (c *UserController) SignUp() {
	isAjax := c.GetString("isajax")

	if isAjax == "1" {
		username := c.GetString("username")
		password := c.GetString("password")
		repass := c.GetString("repass")
		email := c.GetString("email")

		newUser := m.User{
			Username:   username,
			Password:   password,
			Repassword: repass,
			Email:      email,
		}

		_, err := m.AddUser(&newUser)

		if err == nil {
			c.Resp(true, "注册成功！欢迎使用你的个人词典！")
			return
		}
		c.Resp(false, err.Error())
		return

	}
}

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
