package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "hello/routers"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:wandl123@tcp(127.0.0.1:3306)/imooc")
	beego.Run()
}
