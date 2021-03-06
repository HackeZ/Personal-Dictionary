package routers

import (
	"Personal-Dictionary/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")

	/***** User Router Start *****/

	// Login.
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	// Sign Up.
	beego.Router("/signup", &controllers.UserController{}, "*:SignUp")
	// Logout
	beego.Router("/logout", &controllers.UserController{}, "*:Logout")

	/***** User Router End *****/

	// PD
	beego.Router("/pd", &controllers.MainController{}, "*:PdIndex")
	beego.Router("/pd/add", &controllers.MainController{}, "post:AddPersonalDictionary")
	beego.Router("/pd/del", &controllers.MainController{}, "*:DelPersonalDictionary")
}
