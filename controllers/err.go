package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["code"] = 404
	c.TplName = "err.html"
}

func (c *ErrorController) Error500() {
	c.Data["code"] = 500
	c.TplName = "err.html"
}
