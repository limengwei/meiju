package admin

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

///client ip
func (this *baseController) getClientIp() string {
	return this.Ctx.Input.IP()
}

///获取当前时间
func (this *baseController) getTime() time.Time {
	return time.Now()
	//add := 8 * float64(time.Hour)
	//return time.Now().UTC().Add(time.Duration(add))
}

//是否post提交
func (this *baseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//format time
func (this *baseController) FormatTime(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

func (this *baseController) GetGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}
