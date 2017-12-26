package admin

import (
	"github.com/astaxie/beego"
)

type MovieController struct {
	beego.Controller
}

func (c *MovieController) Get() {
	var pageIndex = 1
	pageIndex, _ = c.GetInt("p", 1)

	var movieList = make([]string, 10)
	c.Data["movieList"] = movieList
	c.Data["total"] = 450
	c.Data["p"] = pageIndex
	c.TplName = "admin/movie-list.html"
}

func (c *MovieController) Add() {
	c.TplName = "admin/movie-edit.html"
}

func (c *MovieController) Edit() {
	c.TplName = "admin/movie-edit.html"
}
