package main

import (
	_ "meiju/routers"

	"github.com/astaxie/beego"

	_ "meiju/logger"

	_ "meiju/tools"
)

func main() {
	//防XSRF:跨站请求伪造(暂不启用)
	//	beego.BConfig.WebConfig.EnableXSRF = true
	//	beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	//	beego.BConfig.WebConfig.XSRFExpire = 3600

	beego.Run()
}
