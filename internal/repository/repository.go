package repository

import (
	"log"
	"time"

	"github.com/pavkosykh/gallifrey/config"
	"github.com/pavkosykh/gallifrey/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartTimeSerial(word string) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.SQLITE.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection error: %s", err)
	}

	var lastSearial model.TimeSerial
	db.Last(&lastSearial)
	if !lastSearial.StopAt.Valid {
		db.Model(&lastSearial).Update("StopAt", time.Now().Format(time.RFC3339))
	}

	var event model.Event
	db.FirstOrCreate(&event, model.Event{Title: word})

	db.Create(&model.TimeSerial{Event: event})
}

func StopTimeSerial() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.SQLITE.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection error: %s", err)
	}

	var lastSearial model.TimeSerial
	db.Last(&lastSearial)

	if !lastSearial.StopAt.Valid {
		db.Model(&lastSearial).Update("StopAt", time.Now().Format(time.RFC3339))
	}
}
