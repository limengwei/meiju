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

	c.TplName = "index.html"
}

func (c *MainController) List() {
	cid := c.Ctx.Input.Param(":cid")

	var list = make([]string, 40)
	var randomList = make([]string, 6)

	c.Data["cid"] = cid
	c.Data["list"] = list
	c.Data["randomList"] = randomList
	c.TplName = "list.html"
}

func (c *MainController) News() {

}
