package cli

import (
	"log"

	"github.com/pavkosykh/gallifrey/config"
	"github.com/pavkosykh/gallifrey/internal/cmd"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.Open(sqlite.Open(cfg.SQLITE.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer CloseConnection(db)

	cmd.Execute()
}

func CloseConnection(db *gorm.DB) {
	dbObj, _ := db.DB()
	dbObj.Close()
}
