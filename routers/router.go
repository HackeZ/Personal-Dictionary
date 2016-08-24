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
	beego.Router("/signup", &controllers.UserController{}, "post:SignUp")

	/***** User Router End *****/

	// PD
	beego.Router("/pb", &controllers.MainController{}, "*:PdIndex")
	beego.Router("/pd/add", &controllers.MainController{}, "post:AddPersonalDictionary")
}
