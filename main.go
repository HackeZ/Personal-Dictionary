package main

import (
	"os"
	"strings"

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

// showWrap Show Wrap in Templete.
func showWrap(input string) (output string) {
	output = strings.Replace(input, " ", " ", -1)
	output = strings.Replace(output, "\n", "<br />", -1)
	return
}

func main() {
	initialize()

	beego.AddFuncMap("pd_markdown", markdown)
	beego.AddFuncMap("pd_showwrap", showWrap)
	beego.Run("127.0.0.1:" + beego.AppConfig.String("httpport"))
}
