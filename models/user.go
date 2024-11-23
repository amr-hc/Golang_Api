package models

import (
	"errors"
	"example.com/api/db"
	"example.com/api/utils"
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

	u.Password, err = utils.CreateHash(u.Password)
	
	if err!= nil {
        return err
    }

	result , err := stmt.Exec(u.Email, u.Password)

	if err!= nil {
        return err
    }

	u.ID , err = result.LastInsertId()
	
	return err
}


func (u *User) Login() error {
	statement := `select id, password from users where email = ?`

	stmt, err := db.DB.Prepare(statement)
	if err!= nil {
        return err
    }

	defer stmt.Close()

	row := db.DB.QueryRow(statement, u.Email)

	var hashPassword string

	err = row.Scan(&u.ID, &hashPassword)

	if err!= nil {
        return err
    }

	var ValidPasswod bool = utils.VerifyPassword(hashPassword, u.Password)

	if !ValidPasswod {
        return errors.New("invalid password")
    }

	return nil
}
