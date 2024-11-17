package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_labs?parseTime=true")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable(){
	createUsersTable :=
	`CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}

	createEventsTable :=
	`CREATE TABLE IF NOT EXISTS events (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        date DATETIME NOT NULL,
		location TEXT NOT NULL,
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}