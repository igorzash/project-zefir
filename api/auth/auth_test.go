package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/igorzash/project-zefir/test"
	"github.com/igorzash/project-zefir/userpkg"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type AuthSuite struct {
	test.Suite
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (suite *AuthSuite) TestLoginHandlerUnauthorized() {
	// Create a request to login
	loginJSON := `{"email": "test@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Send the request to the API
	suite.R.ServeHTTP(resp, req)
	suite.Equal(http.StatusUnauthorized, resp.Code)
}

func (suite *AuthSuite) TestLoginHandlerAuthorized() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	suite.NoError(err)

	currentTime := time.Now().Format(time.RFC3339)
	user := userpkg.User{
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
		PasswordHash: string(hashedPassword),
	}
	gofakeit.Struct(&user)
	_, err = suite.Repos.UserRepo.Insert(&user)
	suite.NoError(err)

	// Create a request to login
	loginJSON := fmt.Sprintf(`{"email": "%s", "password": "password"}`, user.Email)
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Send the request to the API
	suite.R.ServeHTTP(resp, req)
	suite.Equal(http.StatusOK, resp.Code)
}
