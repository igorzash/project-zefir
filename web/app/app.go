package app

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/auth"
	"github.com/igorzash/project-zefir/web/db"
	"github.com/igorzash/project-zefir/web/repos"
)

type AppParams struct {
	RunMigrations  bool
	ServeTemplates bool
	ServeStatic    bool
}

type App struct {
	R      *gin.Engine
	Repos  *repos.Repositories
	DBConn *sql.DB
}

func NewApp(params *AppParams) (*App, error) {
	app := &App{}
	app.R = gin.Default()

	var err error
	app.DBConn, err = db.Connect()
	if err != nil {
		return nil, err
	}

	if params.RunMigrations {
		app.runMigrations()
	}

	app.Repos, err = repos.NewRepositories(app.DBConn)
	if err != nil {
		return nil, err
	}

	if params.ServeTemplates {
		app.R.LoadHTMLGlob("templates/*")
	}
	if params.ServeStatic {
		app.R.Static("/static", "./static")
	}

	auth.SetUpRoutes(app.R, app.Repos)

	app.R.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	app.R.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{})
	})

	app.R.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	return app, nil
}
