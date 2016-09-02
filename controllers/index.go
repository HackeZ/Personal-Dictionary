package controllers

import (
	"Personal-Dictionary/utils"
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
	c.Redirect("/pd", 302)
}

// PdIndex Show PdIndex
func (c *MainController) PdIndex() {

	// runmode = dev
	user := c.GetSession("userinfo")
	var pdLen int64

	c.Data["Title"] = beego.AppConfig.String("login_title")

	if c.GetSession("userinfo") != nil {
		user := c.GetSession("userinfo").(string)
		c.Data["Title"] = "欢迎来到 " + user + " 的个人词典."
		c.Data["User"] = user
		c.Data["PersonalDictionary"], pdLen = m.GetPersonalDictionaryList(user, -1, 0, "Createtime")
	}

	if pdLen == 0 {
		c.Data["PDEmpty"] = true
		c.Data["Tips"] = "你还没有创建过属于你的字典，快来试试吧～"
	} else {
		c.Data["PDEmpty"] = false
		c.Data["Tips"] = ""
	}

	log.Println("UserSession --> ", user)
	c.TplName = "index.tpl"
}

// AddPersonalDictionary 添加一条个人词典
func (c *MainController) AddPersonalDictionary() {

	// runmode = dev
	user := c.GetSession("userinfo")
	var userString string
	if user != nil {
		userString = c.GetSession("userinfo").(string)
	} else {
		userString = "HackerZ"
	}

	// runmode = product
	if "dev" != beego.AppConfig.String("runmode") {
		// Login Check.
		if user == nil {
			c.Resp(false, "你还没有登录，请登录后再试！")
			return
		}
	}

	loginUser, _ := m.GetUserByUsername(userString)

	keyword := c.GetString("Keyword")
	content := c.GetString("Content")

	// 数据合法性判断
	if keyword == "" || content == "" {
		c.Resp(false, "请先填写好数据再提交！")
		return
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

func (c *MainController) DelPersonalDictionary() {
	// runmode = dev
	// user := "HackerZ"

	// runmode = product
	if "dev" != beego.AppConfig.String("runmode") {
		// Login Check.
		user := c.GetSession("userinfo").(string)

		if user == "" {
			c.Resp(false, "你还没有登录，请登录后再试！")
			return
		}
	}

	// loginUser, _ := m.GetUserByUsername(user)

	pdIdString := c.GetString("pd_id")
	pdId, _ := utils.Atoi64(pdIdString)

	log.Println("del pd_id -->", pdId)

	// 数据合法性判读

	_, err := m.DelPersonalDictionary(pdId)
	if err != nil {
		c.Resp(false, err.Error())
		return
	}
	c.Resp(true, "删除成功～ 3s后自动刷新页面～")
	return
}
