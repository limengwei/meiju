package admin

import (
	"crypto/tls"
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

type MailController struct {
	beego.Controller
}

func (c *MailController) Send() {

	if c.Ctx.Request.Method == "GET" {

		c.TplName = "mail.html"

		return
	}

	d := gomail.NewDialer("smtp.qq.com", 465, "892642257@qq.com", "aseqjcbjwgjibbjh")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "892642257@qq.com", "limengweiwei")
	m.SetHeader("To", m.FormatAddress("limengweiwei@gmail.com", "史丹利"))
	m.SetHeader("Subject", "hello world")
	m.SetBody("text/html", "<h1>hello world</h1>")

	c.Data["msg"] = "发送成功"

	if err := d.DialAndSend(m); err != nil {
		c.Data["msg"] = fmt.Sprintln("发送失败：", err)
	}

	c.TplName = "mail.html"
}
