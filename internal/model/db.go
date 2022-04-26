package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase(databseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databseUrl), &gorm.Config{})

	if err != nil {
		log.Printf("Troubles with database: %s", err)
	}

	return db, err
}
