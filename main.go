package main

import (
	"log"

	"github.com/gitfong/lau/db"
)

func main() {
	defer db.DB.Close()

	err := initRouter().Run(":9090")
	if err != nil {
		log.Fatal(err.Error())
	}
}
