package controllers

import "github.com/astaxie/beego"

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
	c.Data["Title"] = beego.AppConfig.String("login_title")
	c.Data["Welcome"] = "Welcome to your Personal Dictionary."
	c.TplName = "index.tpl"
}
