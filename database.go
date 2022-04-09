package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

const database = "gallifrey.sqlite3"

func ensure_database() *sql.DB {
	sqliteDatabase, _ := sql.Open("sqlite3", database)
	return sqliteDatabase
}

func run_migrations(db *sql.DB) {
	migration_create_events_table(db)
	migration_create_table_for_time_series(db)
	migration_create_basic_events(db)

	log.Println("All migrations are completed")
}

func drop_database() {
	e := os.Remove(database)
	if e != nil {
		log.Fatal(e)
	}
}

func close_database(db *sql.DB) {
	db.Close()
}
