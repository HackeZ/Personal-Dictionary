package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

type CommonController struct {
    beego.Controller
}

func init()  {
    accessRegister()
}

// AccessRegister 验证登陆权限
func accessRegister()  {
    var Check = func (ctx *context.Context)  {
        uinfo := ctx.Input.Session("userinfo")
        if uinfo == nil {
            // have no user info, relogin
            ctx.Redirect(302,"/login")
        }
    }
}

// Resp Response request Status and Info.
func (c *CommonController) Resp(status bool, str string) {
	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	c.ServeJSON()
}