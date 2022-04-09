package main

func main() {
	db := ensure_database()
	run_migrations(db)
	close_database(db)
}
