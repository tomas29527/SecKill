package main

import (
	"SecKill/conf"
	"github.com/astaxie/beego"
)

func main() {
	//初始化日志
	conf.LogSetting()

	beego.Run()
}
