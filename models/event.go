package models

import "github.com/duyanhitbe/go-basic-rest/database"

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	UserId      int    `json:"user_id"`
}

var events = []Event{}

// Create inserts a new event into the database.
// It sets the UserId field of the event to 1 and executes an SQL INSERT statement.
// Returns an error if there was a problem preparing the statement, executing the statement,
// or retrieving the last inserted ID from the result.
// If successful, it updates the Id field of the event and appends the event to the events slice.
func (event *Event) Create() error {
	event.UserId = 1
	createEventQuery := `
		INSERT INTO events (title, description, location, user_id) 
		VALUES ($1, $2, $3, $4)
	`
	statement, err := database.DB.Prepare(createEventQuery)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(event.Title, event.Description, event.Location, event.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.Id = int(id)
	events = append(events, *event)
	return nil
}

// GetAllEvents retrieves all events from the database.
// It returns a slice of Event objects and an error if any.
func GetAllEvents() ([]Event, error) {
	getAllEventQuery := `
		SELECT * FROM events
	`
	rows, err := database.DB.Query(getAllEventQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// GetOneEventById retrieves a single event from the database based on the provided ID.
// It returns a pointer to the Event struct and an error if any occurred.
func GetOneEventById(id int64) (*Event, error) {
	getOneEventQuery := `
		SELECT * FROM events WHERE id=$1
	`
	row := database.DB.QueryRow(getOneEventQuery, id)
	var event Event
	err := row.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// UpdateEventById updates an event in the database with the specified ID.
// It takes the ID of the event to be updated and a pointer to the updated event object.
// It returns the updated event object and an error if any occurred.
func UpdateEventById(id int64, event *Event) (*Event, error) {
	updateEventQuery := `
		UPDATE events SET title=$1, description=$2, location=$3 WHERE id=$4
	`
	statement, err := database.DB.Prepare(updateEventQuery)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Title, event.Description, event.Location, id)
	if err != nil {
		return nil, err
	}
	return GetOneEventById(id)
}

// DeleteEventById deletes an event from the database by its ID.
// It returns the deleted event and any error encountered.
func DeleteEventById(id int64) (*Event, error) {
	event, err := GetOneEventById(id)
	if err != nil {
		return nil, err
	}
	deleteEventQuery := `
		DELETE FROM events WHERE id=$1
	`
	statement, err := database.DB.Prepare(deleteEventQuery)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}
