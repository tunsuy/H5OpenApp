package controllers

import (
	"github.com/astaxie/beego"
)

type OpenappSchemeController struct {
	beego.Controller
}

func (c *OpenappSchemeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "openapp_scheme.html"
}
