package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	cfg "goRedisDemo/config"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", cfg.MySQLSource)
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
