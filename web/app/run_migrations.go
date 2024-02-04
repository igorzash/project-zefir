package app

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func (app *App) runMigrations() error {
	driver, err := sqlite3.WithInstance(app.DBConn, &sqlite3.Config{})
	if err != nil {
		return err
	}

	var migrationsURL string
	if _, exists := os.LookupEnv("MIGRATIONS_URL"); exists {
		migrationsURL = os.Getenv("MIGRATIONS_URL")
	} else {
		return fmt.Errorf("MIGRATIONS_URL env variable is not set")
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
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
