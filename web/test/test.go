package test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/igorzash/project-zefir/web/app"
	"github.com/igorzash/project-zefir/web/entities/userpkg"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Suite struct {
	suite.Suite
	App *app.App
}

func (suite *Suite) SetupTest() {
	var err error
	suite.App, err = app.NewApp(&app.AppParams{RunMigrations: true})
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *Suite) TearDownTest() {
	suite.App.DBConn.Close()
}

func (suite *Suite) NewUser() *userpkg.User {
	user, err := userpkg.NewUser(gofakeit.Email(), gofakeit.Username(), gofakeit.Password(true, true, true, false, false, 12))
	suite.NoError(err)
	return user
}

func (suite *Suite) NewUserWithPassword(password string) *userpkg.User {
	user, err := userpkg.NewUser(gofakeit.Email(), gofakeit.Username(), password)
	suite.NoError(err)
	return user
}
