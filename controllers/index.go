package controllers

import (
	"log"

	m "Personal-Dictionary/models"

	"github.com/astaxie/beego"
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
	user := c.GetString("userinfo")

	log.Println("UserSession --> ", user)
	c.Data["Title"] = beego.AppConfig.String("login_title")
	c.Data["Welcome"] = "Welcome to " + user + "' Personal Dictionary."
	c.TplName = "index.tpl"
}

// AddPersonalDictionary 添加一条个人词典
func (c *MainController) AddPersonalDictionary() {

	// Login Check.
	user := c.GetSession("userinfo").(string)

	if user == "" {
		c.Resp(false, "你还没有登录，请登录后再试！")
		return
	}
	loginUser, _ := m.GetUserByUsername(user)

	keyword := c.GetString("Keyword")
	content := c.GetString("Content")

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
}
