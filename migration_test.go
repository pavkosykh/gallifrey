package main

import (
	"log"
	"testing"
)

func TestEventsTableExistsAndFilled(t *testing.T) {
	db := ensure_database()
	row, err := db.Query("SELECT COUNT(*) FROM events")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var count int
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&count)
	}

	if count <= 0 {
		t.Errorf("Table Events is empty or doesn't exist")
	}
}
func TestTimeSeriesTableExists(t *testing.T) {
	db := ensure_database()

	_, err := db.Query("SELECT COUNT(*) FROM time_series")
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestEventssTableExists(t *testing.T) {
	db := ensure_database()

	_, err := db.Query("SELECT COUNT(*) FROM events")
	if err != nil {
		t.Errorf(err.Error())
	}

}
