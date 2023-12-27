package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic(err)
	}
	DB = db

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			title TEXT, 
			description TEXT,
			location TEXT,
			user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}
}
