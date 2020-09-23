package main

import (
	"github.com/gitfong/lau/db"
)

func main() {
	defer db.SqlDB.Close()

	initRouter().Run() // listen and serve on 0.0.0.0:8080
}
