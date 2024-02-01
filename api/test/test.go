package test

import (
	"database/sql"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/igorzash/project-zefir/auth"
	"github.com/igorzash/project-zefir/repos"
	"github.com/igorzash/project-zefir/userpkg"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func SetupGin(repos *repos.Repositories) *gin.Engine {
	r := gin.Default()
	authMiddleware := auth.GetMiddleware(repos)
	r.POST("/login", authMiddleware.LoginHandler)

	return r
}

func SetupDatabase() *sql.DB {
	var err error
	dbConn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Run migrations
	driver, err := sqlite3.WithInstance(dbConn, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations/migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	return dbConn
}

type Suite struct {
	suite.Suite
	R      *gin.Engine
	DBConn *sql.DB
	Repos  *repos.Repositories
}

func (suite *Suite) SetupTest() {
	suite.DBConn = SetupDatabase()

	var err error
	suite.Repos, err = repos.NewRepositories(suite.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize repositories: " + err.Error())
	}

	suite.R = SetupGin(suite.Repos)
}

func (suite *Suite) TearDownTest() {
	suite.DBConn.Close()
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
