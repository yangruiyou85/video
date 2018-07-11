package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {

	dbConn, err = sql.Open("mysql", "root:video.COM123@tcp(10.188.12.11:3306)/video?charset=utf8")

	//dbConn, err = sql.Open("mysql", "root:mysql@tcp(localhost:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())

	}
}
