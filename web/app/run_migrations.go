package app

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const DEFAULT_MIGRATIONS_DIR = "file://../../web_migrations/migrations"

func (app *App) runMigrations() error {
	driver, err := sqlite3.WithInstance(app.DBConn, &sqlite3.Config{})
	if err != nil {
		return err
	}

	var migrationsDir string
	if _, exists := os.LookupEnv("MIGRATIONS_DIR"); exists {
		migrationsDir = os.Getenv("MIGRATIONS_DIR")
	} else {
		migrationsDir = DEFAULT_MIGRATIONS_DIR
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		migrationsDir,
		"sqlite3",
		driver,
	)
	fmt.Println(migrations)
	if err != nil {
		return err
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
