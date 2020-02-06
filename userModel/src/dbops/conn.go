package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

//程序启动的时候就会配置
func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(212.64.24.117:3306)/juece_video?charset=utf8&readTimeout=10s")
	if err != nil {
		panic(err.Error())
	}
}
