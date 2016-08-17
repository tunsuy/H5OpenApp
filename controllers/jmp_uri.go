package controllers

import (
	"github.com/astaxie/beego"
)

type JmpUriController struct {
	beego.Controller
}

func (c *JmpUriController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "jmp_uri.html"
}
