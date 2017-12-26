package controllers

import (
	"fmt"

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

	c.Data["title"] = "měi "

	c.Data["bannerList"] = bannerList
	c.Data["newList"] = newList
	c.Data["hotList"] = hotList
	c.Data["lastList"] = lastList
	c.Data["todayList"] = todayList
	c.Data["newsList"] = newsList
	c.Data["cid"] = 0

	c.TplName = "index.html"
}

func (c *MainController) List() {
	var cid = 0
	var pageIndex = 1

	cid, _ = c.GetInt("cid", 0)
	pageIndex, _ = c.GetInt("p", 1)

	var list = make([]string, 40)
	var randomList = make([]string, 6)

	c.Data["cid"] = cid
	c.Data["list"] = list
	c.Data["randomList"] = randomList
	c.Data["total"] = 658
	c.Data["p"] = pageIndex
	c.TplName = "list.html"
}

func (c *MainController) Detail() {

	var id = c.Ctx.Input.Param("id")

	fmt.Println(id)

	var randomList = make([]string, 6)
	c.Data["randomList"] = randomList
	c.Data["cid"] = 0

	c.TplName = "detail.html"
}

func (c *MainController) News() {

	c.TplName = "pin.html"
}
