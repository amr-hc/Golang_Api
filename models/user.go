package models

import (
	"example.com/api/db"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u *User) Signup() error {
	insert := `INSERT INTO users (email, password) VALUES (?,?)`

	stmt, err := db.DB.Prepare(insert)

	if err!= nil {
        return err
    }

	defer stmt.Close()

	result , err := stmt.Exec(u.Email, u.Password)

	if err!= nil {
        return err
    }

	u.ID , err = result.LastInsertId()
	
	return err
}
