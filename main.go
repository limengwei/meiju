package main

import (
	_ "meiju/routers"

	"github.com/astaxie/beego"

	_ "meiju/logger"

	_ "meiju/tools"
)

func main() {
	//todo
	// 过滤器，XSRF配置，登录cookie操作
	beego.Run()
}
