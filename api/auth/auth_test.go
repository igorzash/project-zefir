package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/igorzash/project-zefir/test"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AuthSuite struct {
	test.Suite
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (suite *AuthSuite) TestLoginHandlerUnauthorized() {
	loginJSON := `{"email": "test@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.R.ServeHTTP(resp, req)
	suite.Equal(http.StatusUnauthorized, resp.Code)
}

func (suite *AuthSuite) TestLoginHandlerAuthorized() {
	password := "password"
	user := suite.NewUserWithPassword(password)

	_, err := suite.Repos.UserRepo.Insert(user)
	suite.NoError(err)

	loginJSON := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, user.Email, password)
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.R.ServeHTTP(resp, req)
	suite.Equal(http.StatusOK, resp.Code)
}
