package config

import (
	"log"

	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	var err error

	db, err := gorm.Open(gormlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}
