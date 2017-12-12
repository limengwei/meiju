package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var bannerList = make([]string, 5)

	var newList = make([]string, 8)

	var hotList = make([]string, 18)

	var lastList = make([]string, 18)

	var todayList = make([]string, 8)

	var newsList = make([]string, 10)

	c.Data["title"] = "meiju"

	c.Data["bannerList"] = bannerList
	c.Data["newList"] = newList
	c.Data["hotList"] = hotList
	c.Data["lastList"] = lastList
	c.Data["todayList"] = todayList
	c.Data["newsList"] = newsList

	//	c.TplName = "index.html"
	c.TplName = "upload.html"
}
