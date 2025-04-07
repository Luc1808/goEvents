package models

import (
	"Luc1808/goEvents/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"createdat"`
	UserId      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (title, description, createdat, user_id) 
	VALUES ($1, $2, $3, $4)
	RETURNING id	
	`

	err := db.DB.QueryRow(query, e.Title, e.Description, e.Created, e.UserId).Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Created, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Created, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE events SET title = $1, description = $2, createdat = $3, user_id = $4 WHERE id = $5`

	_, err := db.DB.Exec(query, e.Title, e.Description, e.Created, e.UserId, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = $1`

	_, err := db.DB.Exec(query, e.ID)
	if err != nil {
		return err
	}

	return nil
}
