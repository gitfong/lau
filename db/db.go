package db

import (
	"database/sql"
	"log"
	"time"

	// db driver
	_ "github.com/go-sql-driver/mysql"
)

// SqlDB db
var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/lau") // todo should be configurable
	if err != nil {
		log.Fatal(err.Error())
	}

	SqlDB.SetMaxOpenConns(20) // coreNum * 2 + diskNum
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetConnMaxIdleTime(10 * time.Minute)
	SqlDB.SetConnMaxLifetime(5 * time.Hour)

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
