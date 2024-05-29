package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Localtion   string
	DateTime    time.Time
	userId      int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to a database ...
	events = append(events, e)
}
