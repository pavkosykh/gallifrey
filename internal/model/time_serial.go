package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type TimeSerial struct {
	gorm.Model
	StopAt  sql.NullTime
	EventID int
	Event   Event
}
