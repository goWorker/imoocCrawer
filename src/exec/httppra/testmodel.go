package main

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wandl123@tcp(127.0.0.1:3306)/mytest?charset=utf8", 30)

}
