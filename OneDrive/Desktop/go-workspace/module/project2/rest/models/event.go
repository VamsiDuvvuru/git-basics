package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
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
	//events = append(events, event)
	AddEventToDB(event)
}

func GetAllEvents() []Event {
	events, err := GetAllEventsFromDB()
	if err != nil {
		fmt.Println("Error fetching events from database:", err)
		return nil
	}
	if len(events) == 0 {
		fmt.Println("No events found in database")
		return nil
	}
	return events
}

func DeleteEvent(id int) {
	err := DeleteEventFromDB(id)
	if err != nil {
		fmt.Println("data not found event.ID:", id)
		fmt.Println("Error deleting event from database:", err)
		return
	}
}

func UpdateEvent(event Event) {
	err := UpdateEventInDB(event)
	if err != nil {
		fmt.Println("Error updating event in database:", err)
		return
	}
	fmt.Println("available events are :", events)
}

func GetEventByID(id int) *Event {
	event, err := GetEventByIDFromDB(id)
	if err != nil {
		fmt.Println("Error fetching event from database:", err)
		return nil
	}
	return event
}

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	fmt.Println("Database connection opened.")
	if err != nil {
		fmt.Println("Could not connect to database.")
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTableSQL()
	createUserTableSQL()
}

func createEventTableSQL() {
	createEventTableSQL := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"dateTime" DATETIME NOT NULL,
		"location" TEXT NOT NULL,
		"user_id" INTEGER ,
		FOREIGN KEY (user_id) REFERENCES users(id)
	  );`
	statement, err := DB.Prepare(createEventTableSQL)
	if err != nil {

		panic("could not create events table:" + err.Error())
	}
	statement.Exec()
	fmt.Println("Events table created or already exists.")
}

func createUserTableSQL() {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL,
		"password" TEXT NOT NULL
	  );`
	statement, err := DB.Prepare(createUserTableSQL)
	if err != nil {
		panic("could not create users table:" + err.Error())
	}
	statement.Exec()
	fmt.Println("Users table created or already exists.")
}

func AddEventToDB(event Event) error {
	insertEventSQL := `INSERT INTO events (title, description, dateTime, location, user_id) VALUES (?, ?, ?, ?, ?)`
	statement, err := DB.Prepare(insertEventSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.Title, event.Description, event.Date, event.Location, 1)
	if err != nil {
		return err
	}
	fmt.Println("Event added to database successfully:", event.Title)
	return nil
}

func GetAllEventsFromDB() ([]Event, error) {
	rows, err := DB.Query("SELECT id, title, description, dateTime, location FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func DeleteEventFromDB(eventID int) error {
	deleteEventSQL := `DELETE FROM events WHERE id = ?`
	statement, err := DB.Prepare(deleteEventSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(eventID)
	if err != nil {
		return err
	}
	fmt.Println("Event deleted from database successfully, ID:", eventID)
	return nil
}

func UpdateEventInDB(event Event) error {
	updateEventSQL := `UPDATE events SET title = ?, description = ?, dateTime = ?, location = ? WHERE id = ?`
	statement, err := DB.Prepare(updateEventSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.Title, event.Description, event.Date, event.Location, event.ID)
	if err != nil {
		return err
	}
	fmt.Println("Event updated in database successfully, ID:", event.ID)
	return nil
}

func GetEventByIDFromDB(eventID int) (*Event, error) {
	queryEventSQL := `SELECT id, title, description, dateTime, location FROM events WHERE id = ?`
	row := DB.QueryRow(queryEventSQL, eventID)
	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func ConvertStringToTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ConvertTimeToString(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
