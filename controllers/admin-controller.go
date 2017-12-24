package controllers

import (
	"encoding/base64"
	"fmt"

	"github.com/astaxie/beego/context"

	"meiju/tools"

	"github.com/astaxie/beego"
)

var (
	cookie_secret_key []byte = []byte("f6af32fa256195f33e323448ae3c5940")
	cookie_name              = "cookieee"
	cookie_maxAge            = 3600 * 24 * 7
)

type AdminController struct {
	beego.Controller
}

var FilterUser = func(ctx *context.Context) {
	cookieStr := ctx.GetCookie(cookie_name)

	if cookieStr == "" {
		ctx.Redirect(302, "/login")
		return
	}

	fmt.Println(cookieStr)
	cookieBytes, err := base64.StdEncoding.DecodeString(cookieStr)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(302, "/login")
	}
	origData, err := tools.AesDecrypt(cookieBytes, cookie_secret_key)
	if err != nil {
		ctx.Redirect(302, "/login")
	}

	str := string(origData)
	fmt.Println(str)

	if str != "lmw" {
		ctx.Redirect(302, "/login")
	}

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
		fmt.Println("--uanme--", uname)
		if uname == "lmw" {
			result, err := tools.AesEncrypt([]byte(uname), cookie_secret_key)
			if err != nil {
				c.Redirect("/login", 302)
			}
			cookieStr := base64.StdEncoding.EncodeToString(result)
			c.Ctx.SetCookie(cookie_name, cookieStr, cookie_maxAge)
			c.Redirect("/a", 302)
			return
		}

		c.Redirect("/login", 302)

	}

}
