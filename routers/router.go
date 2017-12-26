package routers

import (
	"meiju/controllers"
	"meiju/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{})

	beego.Router("/list", &controllers.MainController{}, "*:List")

	beego.Router("/d/:id", &controllers.MainController{}, "*:Detail")

	beego.Router("/news", &controllers.MainController{}, "*:News")

	beego.InsertFilter("/a/*", beego.BeforeRouter, admin.FilterUser, true, true)

	////admin
	beego.Router("/login", &admin.AdminController{}, "*:Login")

	beego.Router("/logout", &admin.AdminController{}, "*:Logout")

	beego.Router("/mail", &admin.MailController{}, "*:Send")

	ns := beego.NewNamespace("/a",
		beego.NSRouter("/", &admin.AdminController{}),
		beego.NSRouter("/home", &admin.AdminController{}, "*:Home"),
		beego.NSRouter("/movie", &admin.MovieController{}),
		beego.NSRouter("/movie/edit", &admin.MovieController{}, "*:Add"),
		beego.NSRouter("/movie/edit/:id", &admin.MovieController{}, "*:Edit"),

		beego.NSRouter("/relation", &admin.RelationController{}),
		beego.NSRouter("/relation/edit", &admin.RelationController{}, "*:Add"),
		beego.NSRouter("/relation/edit/:id", &admin.RelationController{}, "*:Edit"),
	)

	beego.AddNamespace(ns)
}
