package routers

import (
	"meiju/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{})

	beego.Router("/list", &controllers.MainController{}, "*:List")

	beego.Router("/d/:id", &controllers.MainController{}, "*:Detail")

	beego.Router("/news", &controllers.MainController{}, "*:News")

	beego.Router("/login", &controllers.AdminController{}, "*:Login")
	beego.InsertFilter("/a/*", beego.BeforeRouter, controllers.FilterUser, true, true)

	////admin
	ns := beego.NewNamespace("/a",
		beego.NSRouter("/", &controllers.AdminController{}),
		beego.NSRouter("/movie", &controllers.MovieController{}),
		beego.NSRouter("/movie/edit", &controllers.MovieController{}, "*:Add"),
		beego.NSRouter("/movie/edit/:id", &controllers.MovieController{}, "*:Edit"),
	)

	beego.AddNamespace(ns)
}
