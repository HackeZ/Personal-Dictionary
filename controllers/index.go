package controllers

import (
	"log"

	"github.com/astaxie/beego"

	m "Personal-Dictionary/models"
)

// MainController Personal-Dictionary main Controllers
type MainController struct {
	CommonController
}

// Index Personal Dictionary Index
func (c *MainController) Index() {
	c.Redirect("/pb", 302)
}

// PdIndex Show PdIndex
func (c *MainController) PdIndex() {

	// runmode = dev
	user := c.GetSession("userinfo")

	c.Data["Title"] = beego.AppConfig.String("login_title")
	c.Data["User"] = "HackerZ"

	// runmode = product
	if c.GetSession("userinfo") != nil {
		user := c.GetSession("userinfo").(string)
		c.Data["Title"] = "Welcome to " + user + "' Personal Dictionary."
		c.Data["User"] = user
	}

	log.Println("UserSession --> ", user)
	c.TplName = "index.tpl"
}

// AddPersonalDictionary 添加一条个人词典
func (c *MainController) AddPersonalDictionary() {

	// runmode = dev
	user := "HackerZ"

	// runmode = product
	if "dev" != beego.AppConfig.String("runmode") {
		// Login Check.
		user := c.GetSession("userinfo").(string)

		if user == "" {
			c.Resp(false, "你还没有登录，请登录后再试！")
			return
		}
	}

	loginUser, _ := m.GetUserByUsername(user)

	keyword := c.GetString("Keyword")
	content := c.GetString("Content")

	// 数据合法性判断
	if keyword == "" || content == "" {
		c.Resp(false, "请先填写好数据再提交！")
	}

	pd := new(m.PersonalDictionary)
	pd.User = &loginUser
	pd.Keyword = keyword
	pd.Content = content

	_, err := m.AddPersonalDictionary(pd)
	if err != nil {
		c.Resp(false, err.Error())
		return
	}
	c.Resp(true, "新建词典成功！")
	return
}
