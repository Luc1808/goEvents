package models

import (
	"time"
)

type Event struct {
	id          int
	title       string
	description string
	created     time.Time
	userId      int
}
