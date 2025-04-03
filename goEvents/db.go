package main

import "database/sql"

func StartDB() {
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=event-db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS events (
				id SERIAL PRIMARY KEY,
				title VARCHAR(45),
				description TEXT,
				createdAt timestamp,
				user_id integer
			);
		`)
	if err != nil {
		panic(err)
	}
}
