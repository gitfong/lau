package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	// db driver
	_ "github.com/go-sql-driver/mysql"
)

// DB db
var DB *sql.DB

func init() {
	cfg := loadMySQLCfg()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Account, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	DB.SetMaxOpenConns(20) // coreNum * 2 + diskNum
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(5 * time.Hour)

	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// MySQLCfg configure of mysql
type MySQLCfg struct {
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	DBName   string `json:"db_name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func loadMySQLCfg() *MySQLCfg {
	/*
		// host
		host := os.Getenv("MySQLCfg_host")
		if host == "" {
			log.Fatal("wrong environment variable. MySQLCfg_host is null")
		}

		// port
		portStr := os.Getenv("MySQLCfg_port")
		if portStr == "" {
			log.Fatal("wrong environment variable. MySQLCfg_port is null")
		}
		port, err := strconv.ParseUint(portStr, 10, 32)
		if err != nil {
			log.Fatal(err.Error())
		}

		// database name
		dbName := os.Getenv("MySQLCfg_dbname")
		if dbName == "" {
			log.Fatal("wrong environment variable. MySQLCfg_dbname is null")
		}

		// account
		account := os.Getenv("MySQLCfg_account")
		if account == "" {
			log.Fatal("wrong environment variable. MySQLCfg_account is null")
		}

		// password
		password := os.Getenv("MySQLCfg_password")
		if password == "" {
			log.Fatal("wrong environment variable. MySQLCfg_password is null")
		}

		cfg := MySQLCfg{
			Host:   host,
			Port:   uint32(port),
			DBName: dbName,
			Account: account,
			Password: password,
		}
	*/

	cfg := MySQLCfg{}
	data, err := ioutil.ReadFile("./configs/mysql.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("mysql config:", cfg)

	return &cfg
}
