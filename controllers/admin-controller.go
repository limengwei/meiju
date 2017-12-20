package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {

	c.TplName = "admin/index.html"
}

func (c *AdminController) Login() {

	if c.Ctx.Request.Method == "GET" {
		c.TplName = "admin/login.html"
		return
	}

	if c.Ctx.Request.Method == "POST" {
		uname := c.GetString("uname", "")
		//		pwd := c.GetString("pwd", "")

		fmt.Println("--uanme--", uname)

		c.Redirect("/a/login", 301)
	}

}
