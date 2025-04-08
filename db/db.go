package db

import "database/sql"

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=event-db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Could not start DB connection.")
	}
	// defer DB.Close()

	createTables()
}

func createTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			email VARCHAR(50) NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		panic("Could not create the DB table users.")
	}

	_, err = DB.Exec(`
			CREATE TABLE IF NOT EXISTS events (
				id SERIAL PRIMARY KEY,
				title VARCHAR(45),
				description TEXT,
				createdAt TIMESTAMP NOT NULL,
				user_id INTEGER NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`)
	if err != nil {
		panic("Could not create the DB table events.")
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS registrations (
			id SERIAL PRIMARY KEY,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users (id),
			FOREIGN KEY (event_id) REFERENCES events (id) 	
		)	
	`)
	if err != nil {
		panic("Could not create the DB table registrations.")
	}
}
