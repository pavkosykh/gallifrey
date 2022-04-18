package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase(databseUrl string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(databseUrl), &gorm.Config{})

	return
}
