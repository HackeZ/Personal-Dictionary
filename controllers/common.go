package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// CommonController Before every Router.
type CommonController struct {
	beego.Controller
}

func init() {
	if "dev" != beego.AppConfig.String("runmode") {
		var FiltAdmin = func(ctx *context.Context) {
			uinfo := ctx.Input.CruSession.Get("userinfo")
			log.Println("uinfo", uinfo)
			if uinfo == nil {
				ctx.Redirect(302, "/login")
			}
		}
		// 在访问 "/pb/*" URI 时，寻找路由器之前(BeforeRouter)，先进行 Check 验证访问过滤器
		beego.InsertFilter("/pd", beego.BeforeRouter, FiltAdmin)
	}
}

// Resp Response request Status and Info.
func (c *CommonController) Resp(status bool, str string) {
	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	c.ServeJSON()
}

// InfoJump Jump to Info Page and Show Info.
func (c *CommonController) InfoJump(info, jumpToPath string) {
	c.Data["Info"] = info
	c.Data["JumpTo"] = jumpToPath
	c.Data["JTS"], _ = beego.AppConfig.Int("info_jump_time")
	c.TplName = "info.tpl"
}
