package main

import (
	"os"

	m "Personal-Dictionary/models"
	_ "Personal-Dictionary/routers"
	"Personal-Dictionary/utils"

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

// Support Markdown Function.
func markdown(input string) (output string) {
	output = utils.StringsToMarkdown(input)
	return
}

func main() {
	initialize()

	beego.AddFuncMap("pd_markdown", markdown)
	beego.Run("127.0.0.1:" + beego.AppConfig.String("httpport"))
}
