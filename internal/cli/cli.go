package cli

import (
	"log"

	"github.com/pavkosykh/gallifrey/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.Open(sqlite.Open(cfg.SQLITE.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer CloseConnection(db)
}

func CloseConnection(db *gorm.DB) {
	dbObj, _ := db.DB()
	dbObj.Close()
}
