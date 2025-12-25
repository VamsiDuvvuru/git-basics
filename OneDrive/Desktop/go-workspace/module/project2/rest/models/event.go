package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
}

func NewEvent(id int, title string, description string, date time.Time, location string) *Event {
	return &Event{
		ID:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Location:    location,
	}
}

var events []Event = []Event{}

func AddEvent(event Event) {
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}

func DeleteEvent(event Event) {
	for i, e := range events {
		if e.ID == event.ID {
			events = append(events[:i], events[i+1:]...)
			break
		}
	}
}

func UpdateEvent(event Event) {
	for i, e := range events {
		if e.ID == event.ID {
			events[i] = event
			break
		}
	}
	//return error if event not found
	panic(fmt.Sprintf("data not found event.ID: %d", event.ID))
}

func GetEventByID(id int) *Event {
	for _, e := range events {
		if e.ID == id {
			return &e
		}
	}
	return nil
}
