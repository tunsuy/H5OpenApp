package routers

import (
	"github.com/astaxie/beego"
	"tsH5/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/openapp_packet", &controllers.OpenappPacketController{})
	beego.Router("/openapp_scheme", &controllers.OpenappSchemeController{})
	beego.Router("/jmp_uri", &controllers.JmpUriController{})
}
