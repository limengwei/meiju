package logger

import "github.com/astaxie/beego/logs"

func init() {
	log := logs.NewLogger()

	log.SetLogger(logs.AdapterFile, `{"filename":"logger.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async() //异步输出 提升性能
	log.Debug("this is a debug message")
}
