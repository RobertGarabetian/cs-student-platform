package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Open the SQLite database file (it will be created if it doesn't exist)
	DB, err = sql.Open("sqlite3", "./cs_platform.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Verify the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to the SQLite database!")

	// Create tables if they don't exist
	createTables()
}

func createTables() {
	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`
	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	// Create other tables as needed
}
