package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:wandl123@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(10)
}
