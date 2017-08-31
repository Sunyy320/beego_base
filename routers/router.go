package routers

import (
	"mytest/controllers"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	// namespace 使用
	ns:=beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
			return true
		}),
		beego.NSRouter("/test",&controllers.MainController{},"get:Test"),
		beego.NSRouter("/validation",&controllers.MainController{},"post:ValidationTest"),
		beego.NSRouter("/dbinsert",&controllers.MainController{},"post:DbInserTest"),
	)
	beego.AddNamespace(ns)
}
