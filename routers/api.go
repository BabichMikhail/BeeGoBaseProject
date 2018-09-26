package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	NSApi := beego.NewNamespace("/api",
		beego.NSNamespace("/game",
			beego.NSRouter("get", &controllers.GameApiController{}, "get:Get"),
		),
	)
	beego.AddNamespace(NSApi)
}
