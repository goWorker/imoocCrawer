package main

import (
	"github.com/astaxie/beego"
	_ "myblogweb/routers"
	"myblogweb/utils"
)

func main() {
	beego.SetStaticPath("/static", "./static")
	utils.InitMysql()
	beego.Run()
}
