package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectToDb() {
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

	fmt.Println("Connected to SQLite database")
}
