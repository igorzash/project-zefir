package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SqliteConnect() (*sql.DB, error) {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		if _, exists := os.LookupEnv("ALLOW_IN_MEMORY_DB"); exists {
			log.Println("[WARNING] No SQLITE_DB_PATH environment variable found, using in-memory database")
			dbPath = ":memory:"
		} else {
			return nil, fmt.Errorf("SQLITE_DB_PATH environment variable is not set")
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
