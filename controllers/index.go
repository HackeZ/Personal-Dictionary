package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	CommonController
}

// Index Personal Dictionary Index
func (c *MainController) Index() {
	c.Data["welcome"] = "Welcome to your Personal Dictionary."
	c.TplName = "index.tpl"
}
