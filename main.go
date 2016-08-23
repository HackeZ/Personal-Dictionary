package main

import (
	_ "Personal-Dictionary/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run("127.0.0.1:" + beego.AppConfig.String("httpport"))
}
