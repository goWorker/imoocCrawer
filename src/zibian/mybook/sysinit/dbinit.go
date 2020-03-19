package sysinit

import (
	_"zibian/mybook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string){
	isDev := ("dev"==beego.AppConfig.String("runmode"))
	if len(aliases) >0 {
		for _, alias := range aliases{
			registDatabase(alias)
			if "w" == alias{
				orm.RunSyncdb("default",false,true)
			}
		}

	}else{
		registDatabase("w")
		orm.RunSyncdb("default",false,true)
	}

	if isDev {
		orm.Debug = isDev
	}
}

func registDatabase(alias string){
	if len(alias) <= 0 {
		return
	}
	dbAlias := alias
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}
	dbName := beego.AppConfig.String("db_"+alias+"_databse")

	dbuser := beego.AppConfig.String("db_"+alias+"_username")

	dbPwd := beego.AppConfig.String("db_"+alias+"_password")

	dbHost := beego.AppConfig.String("db_"+alias+"_host")

	dbPort := beego.AppConfig.String("db_"+alias+"_port")

	orm.RegisterDataBase(dbAlias,"mysql",dbuser+":"+dbPwd+"@tpc("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8",30)
}