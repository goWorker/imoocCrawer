package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type Page struct{
	Id	int
	Website string
	Email string
}
//root:wandl123@tcp(127.0.0.1:3306)/myblogweb

func init(){
	orm.RegisterDataBase("default","mysql","root:wandl123@tcp(127.0.0.1:3306)/test")
	orm.RegisterModel(new(Page))
}

func GetPage() Page{
//	rtn := Page{Website:"hellobegeo",Email:"allanyang@juniper.net"}
//	return rtn
	o := orm.NewOrm()
	p := Page{Id: 1}
	err := o.Read(&p)
	fmt.Println(err)
	return p
}

func UpdatePage(){
	p := Page{Id:1,Email:"emailupdate"}
	o := orm.NewOrm()
	o.Update(&p,"Email")

}