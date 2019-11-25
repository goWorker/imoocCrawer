package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql")

var (
	dbConn *sql.DB
	err error
)

func init() {

	dbConn, _ = sql.Open("mysql","root:wandl123@tcp(127.0.0.1:3306)/video_server")


	if err != nil {
		panic(err.Error())
	}
}