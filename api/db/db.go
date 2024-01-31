package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func Connect() {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		log.Fatal("SQLITE_DB_PATH environment variable is not set")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	Conn = db

	fmt.Println("Connected to SQLite database")
}
