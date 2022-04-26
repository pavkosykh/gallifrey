package model

import (
	"time"

	"gorm.io/gorm"
)

type TimeSerial struct {
	gorm.Model
	StopAt  time.Time
	EventID int
	Event   Event
}
