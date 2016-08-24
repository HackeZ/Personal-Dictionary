package controllers

import (
	"Personal-Dictionary/utils"
	"errors"
	"log"
	"time"

	m "Personal-Dictionary/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

// UserController Login & Logout & SignUp.
type UserController struct {
	beego.Controller
}

var cpt *captcha.Captcha

// Login do Login or Show Login Page.
func (c *UserController) Login() {

	isAjax := c.GetString("isajax")
	log.Println("isAjax -->", isAjax)

	if isAjax == "1" {
		// Captcha Verification.
		if !cpt.VerifyReq(c.Ctx.Request) {
			c.Resp(false, "验证码错误！")
			return
		}

		// Passing Captcha Verify.
		username := c.GetString("username")
		password := c.GetString("password")

		user, err := doLogin(username, password)
		if err == nil {
			c.SetSession("userinfo", user.Username)
			c.Resp(true, "登陆成功")
			return
		}
		c.Resp(false, err.Error())
		return
	}
	// Login Fail! relogin.
	c.Data["Title"] = beego.AppConfig.String("login_title")

	// Get Verification Code.
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums, _ = beego.AppConfig.Int("captcha_length")
	cpt.StdWidth = 100
	cpt.StdHeight = 42

	c.TplName = "login.tpl"
}

func doLogin(username, password string) (m.User, error) {
	user, _ := m.GetUserByUsername(username)

	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != utils.PassEncode(password, user.Salt) {
		return m.User{}, errors.New("密码错误")
	}
	// 更新最后登录时间
	go updateLastLogintime(user)
	// 清除敏感数据
	user.Salt, user.Password, user.Repassword = "", "", ""

	return user, nil
}

// updateLastLogintime 更新最后登录时间
func updateLastLogintime(user m.User) {
	user.Lastlogintime = time.Now()
	m.UpdateUser(&user)
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
