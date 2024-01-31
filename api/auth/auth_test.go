package auth

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/igorzash/project-zefir/db"
	"github.com/igorzash/project-zefir/test"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandlerUnauthorized(t *testing.T) {
	test.SetupEnvironment()
	defer db.Conn.Close()

	// Initialize your routes
	r := gin.Default()
	authMiddleware := GetMiddleware() // use the imported package to call the function
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
	test.SetupEnvironment()
	defer db.Conn.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	currentTime := time.Now().Format(time.RFC3339)

	_, err = db.Conn.Exec("INSERT INTO users (created_at, updated_at, nickname, email, password_hash) VALUES (?, ?, ?, ?, ?)", currentTime, currentTime, "test", "test@example.com", hashedPassword)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	// Initialize your routes
	r := gin.Default()
	authMiddleware := GetMiddleware() // use the imported package to call the function
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
