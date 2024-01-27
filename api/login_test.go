package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandlerUnauthorized(t *testing.T) {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer DB.Close()

	// Run migrations
	driver, err := sqlite3.WithInstance(DB, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/migrations", // replace with the path to your migrations
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Initialize your routes
	r := gin.Default()
	authMiddleware := SetUpAuthMiddleware() // use the imported package to call the function
	r.POST("/login", authMiddleware.LoginHandler)

	// Create a request to login
	loginJSON := `{"email": "test@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Send the request to the API
	r.ServeHTTP(resp, req)

	// Check the response
	if resp.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401; got %d", resp.Code)
	}

	// You can also check the response body with resp.Body.String()
}

func TestLoginHandlerAuthorized(t *testing.T) {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer DB.Close()

	// Run migrations
	driver, err := sqlite3.WithInstance(DB, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/migrations", // replace with the path to your migrations
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	currentTime := time.Now().Format(time.RFC3339)

	_, err = DB.Exec("INSERT INTO users (created_at, updated_at, nickname, email, password_hash) VALUES (?, ?, ?, ?, ?)", currentTime, currentTime, "test", "test@example.com", hashedPassword)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	// Initialize your routes
	r := gin.Default()
	authMiddleware := SetUpAuthMiddleware() // use the imported package to call the function
	r.POST("/login", authMiddleware.LoginHandler)

	// Create a request to login
	loginJSON := `{"email": "test@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Send the request to the API
	r.ServeHTTP(resp, req)

	// Check the response
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200; got %d", resp.Code)
	}

	// You can also check the response body with resp.Body.String()
}
