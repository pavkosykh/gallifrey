package main

import (
	"os"
	"testing"
)

func TestDatabaseFileExists(t *testing.T) {
	db := ensure_database()
	_, err := os.Stat("gallifrey.sqlite3")
	if err != nil {
		t.Errorf("Database does not exists")
	}
	close_database(db)
}
