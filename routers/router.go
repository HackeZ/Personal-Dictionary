package routers

import (
	"Personal-Dictionary/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")

	// Login
	beego.Router("/login", &controllers.UserController{}, "*:Login")

	// Sign up
	beego.Router("/signup", &controllers.UserController{}, "post:SignUp")

	// PD
	beego.Router("/pb", &controllers.MainController{}, "*:PdIndex")
}
