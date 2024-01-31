package test

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/igorzash/project-zefir/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func SetupEnvironment() {
	var err error
	db.Conn, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Run migrations
	driver, err := sqlite3.WithInstance(db.Conn, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/migrations", // replace with the path to your migrations
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}
