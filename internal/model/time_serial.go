package model

import "gorm.io/gorm"

type TimeSerial struct {
	gorm.Model
	EventID int
	Event   Event
}
