package repository

// import (
// 	"database/sql"
// 	"example/my-project-go/module/project2/rest/models"
// 	"fmt"
// 	"time"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDB() {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "api.db")

// 	if err != nil {
// 		panic("Could not connect to database.")
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	createTables()
// }

// func createTables() {
// 	createEventTableSQL := `CREATE TABLE IF NOT EXISTS events (
// 		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 		"title" TEXT NOT NULL,
// 		"description" TEXT NOT NULL,
// 		"dateTime" DATETIME NOT NULL,
// 		"location" TEXT NOT NULL,
// 		"user_id" INTEGER NOT NULL
// 	  );`
// 	statement, err := DB.Prepare(createEventTableSQL)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	statement.Exec()
// 	fmt.Println("Events table created or already exists.")
// }

// func AddEventToDB(event models.Event) error {
// 	insertEventSQL := `INSERT INTO events (title, description, dateTime, location, user_id) VALUES (?, ?, ?, ?, ?)`
// 	statement, err := DB.Prepare(insertEventSQL)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = statement.Exec(event.Title, event.Description, event.Date, event.Location, 1)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Event added to database successfully:", event.Title)
// 	return nil
// }

// func GetAllEventsFromDB() ([]models.Event, error) {
// 	rows, err := DB.Query("SELECT id, title, description, dateTime, location FROM events")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var events []models.Event
// 	for rows.Next() {
// 		var event models.Event
// 		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location)
// 		if err != nil {
// 			return nil, err
// 		}
// 		events = append(events, event)
// 	}
// 	return events, nil
// }

// func DeleteEventFromDB(eventID int) error {
// 	deleteEventSQL := `DELETE FROM events WHERE id = ?`
// 	statement, err := DB.Prepare(deleteEventSQL)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = statement.Exec(eventID)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Event deleted from database successfully, ID:", eventID)
// 	return nil
// }

// func UpdateEventInDB(event models.Event) error {
// 	updateEventSQL := `UPDATE events SET title = ?, description = ?, dateTime = ?, location = ? WHERE id = ?`
// 	statement, err := DB.Prepare(updateEventSQL)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = statement.Exec(event.Title, event.Description, event.Date, event.Location, event.ID)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Event updated in database successfully, ID:", event.ID)
// 	return nil
// }

// func GetEventByIDFromDB(eventID int) (*models.Event, error) {
// 	queryEventSQL := `SELECT id, title, description, dateTime, location FROM events WHERE id = ?`
// 	row := DB.QueryRow(queryEventSQL, eventID)
// 	var event models.Event
// 	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &event, nil
// }

// func ConvertStringToTime(dateTimeStr string) (time.Time, error) {
// 	layout := "2006-01-02 15:04:05"
// 	parsedTime, err := time.Parse(layout, dateTimeStr)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	return parsedTime, nil
// }

// func ConvertTimeToString(t time.Time) string {
// 	layout := "2006-01-02 15:04:05"
// 	return t.Format(layout)
// }
