package controllers

import (
	"github.com/astaxie/beego"
)

type MovieController struct {
	beego.Controller
}

func (c *MovieController) Get() {
	c.TplName = "admin/movie-list.html"
}

func (c *MovieController) Add() {
	c.TplName = "admin/movie-edit.html"
}

func (c *MovieController) Edit() {
	c.TplName = "admin/movie-edit.html"
}
