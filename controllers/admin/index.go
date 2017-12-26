package admin

import (
	"encoding/base64"
	"fmt"

	"html/template"

	"github.com/astaxie/beego/context"

	"meiju/tools"
	"os"
	"runtime"
	"time"

	"github.com/astaxie/beego"
)

var (
	cookie_secret_key []byte = []byte("f6af32fa256195f33e323448ae3c5940")
	cookie_name              = "cookieee"
	cookie_maxAge            = 3600 * 24 * 7
)

type AdminController struct {
	baseController
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
func (c *AdminController) Home() {
	c.Data["hostname"], _ = os.Hostname()
	c.Data["gover"] = runtime.Version()
	c.Data["os"] = runtime.GOOS
	c.Data["cpunum"] = runtime.NumCPU()
	c.Data["arch"] = runtime.GOARCH
	c.Data["beegover"] = beego.VERSION
	c.Data["clientip"] = c.getClientIp()
	c.Data["systemver"] = beego.AppConfig.String("systemver")
	c.Data["developer"] = beego.AppConfig.String("developer")
	c.Data["servertime"] = c.FormatTime(time.Now(), "YYYY年MM月DD日 HH:mm:ss")
	//影片数量
	//	var movie models.MovieInfo
	//	count, _ := movie.Query().Count()
	c.Data["moviecount"] = 9998
	c.TplName = "admin/home.html"
}

func (c *AdminController) Login() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "admin/login.html"
		return
	}

	if c.Ctx.Request.Method == "POST" {
		uname := c.GetString("uname", "")
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

func (c *AdminController) Logout() {
	c.Ctx.SetCookie(cookie_name, "", 0)
	c.Ctx.Redirect(302, "/login")
}
