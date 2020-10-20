package models

import (
	"github.com/gitfong/lau/db"
)

const (
	sqlLoadInfo = "select name,city,province from author limit 1"
)

// Author author's information
type Author struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Province string `json:"province"`
}

// LoadInfo load author's information from databases
func (a *Author) LoadInfo() error {
	rows, err := db.DB.Query(sqlLoadInfo)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&a.Name, &a.City, &a.Province)
	}

	return nil
}
