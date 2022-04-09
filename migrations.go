package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func migration_create_events_table(db *sql.DB) {
	SQL := `CREATE TABLE IF NOT EXISTS events (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"title" TEXT,
			UNIQUE(title)

		);`
	execute(db, SQL, "Events table is created")
}

func migration_create_table_for_time_series(db *sql.DB) {
	SQL := `CREATE TABLE IF NOT EXISTS time_series (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"start_datetime" DATETIME,
		"end_datetime" DATETIME,
		"event_id" INTEGER,
		FOREIGN KEY ("event_id") REFERENCES events(id)
	  );`
	execute(db, SQL, "Time Series table is created")

}

func migration_create_basic_events(db *sql.DB) {
	SQL := `INSERT OR IGNORE INTO events (title)
		VALUES ("food"), ("sport"), ("job"), ("games") 
	  `
	execute(db, SQL, "Events table is filled with basic values")

}

func execute(db *sql.DB, sql string, log_message string) {
	statement, err := db.Prepare(sql) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println(log_message)
}
