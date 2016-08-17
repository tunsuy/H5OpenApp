package controllers

import (
	"github.com/astaxie/beego"
)

type OpenappPacketController struct {
	beego.Controller
}

func (c *OpenappPacketController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "openapp_packet.html"
}
