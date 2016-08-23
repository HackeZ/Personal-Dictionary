package routers

import (
	"Personal-Dictionary/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")

	// Login
	beego.Router("/login", &controllers.UserController{}, "*:Login")

	// PD
	beego.Router("/pb", &controllers.MainController{}, "*:PdIndex")
}
