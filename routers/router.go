package routers

import (
	"meiju/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/list", &controllers.MainController{}, "*:List")

	beego.Router("/news", &controllers.MainController{}, "*:News")
}
