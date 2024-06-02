package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding: required`
	Description string    `binding: required`
	Localtion   string    `binding: required`
	DateTime    time.Time `binding: required`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
