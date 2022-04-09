package main

// Prepare before ALL tests
func init() {
	// Here we drop old database and create it from the start
	drop_database()
	db := ensure_database()
	run_migrations(db)
	close_database(db)
}
