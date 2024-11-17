package models

import (
	"time"
	"example.com/api/db"
)

type Event struct {
	ID          int64     
	Title       string  `binding:"required"`
	Description string  `binding:"required"`
	Date        time.Time  `binding:"required"`
	Location    string  `binding:"required"`
	UserID      int   
}


func (e *Event) Save() error {
	insert := `INSERT INTO events (name, description, date, location , user_id) VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(insert)

	if err!= nil {
        return err
    }

	defer stmt.Close()

	result , err := stmt.Exec(e.Title, e.Description, e.Date, e.Location, e.UserID)

	if err!= nil {
        return err
    }

	e.ID , err = result.LastInsertId()
	
	return err
}

func GetAllEvents() ([]Event, error) {
	statement := "SELECT * FROM events"
	rows, err := db.DB.Query(statement)
	if err != nil {
        return nil , err
    }
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
		if err != nil{
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}



func GetEventById(id int64) (*Event, error) {
	statement := "SELECT * FROM events where id = ?"
	row := db.DB.QueryRow(statement, id)
	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
	if err != nil{
		return nil, err
	}
	return &event, nil
}

func (e *Event) Update() error {
	update := `
	UPDATE events
	SET name = ?, description = ?, date = ?, location = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(update)

	if err != nil {
        return err
    }

	defer stmt.Close()

	_ , err = stmt.Exec(e.Title, e.Description, e.Date, e.Location, e.ID)

	return err
}

func (e *Event) Delete() error {
	statement := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(statement)

	if err!= nil {
        return err
    }

	defer stmt.Close()

	_ , err = stmt.Exec(e.ID)

	return err
}