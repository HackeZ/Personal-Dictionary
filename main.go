package main

import (
	"os"

	m "Personal-Dictionary/models"
	_ "Personal-Dictionary/routers"

	"github.com/astaxie/beego"
)

func initialize() {

	initArgs()

	m.Connect()
}

func initArgs() {
	args := os.Args

	for _, v := range args {
		if v == "-syncdb" {
			m.Syncdb()
			os.Exit(0)
		}
	}
}

func main() {
	initialize()

	beego.Run("127.0.0.1:" + beego.AppConfig.String("httpport"))
}
