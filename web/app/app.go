package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/auth"
	"github.com/igorzash/project-zefir/web/controllers"
	"github.com/igorzash/project-zefir/web/db"
	"github.com/igorzash/project-zefir/web/entities"
)

type AppParams struct {
	RunMigrations  bool
	ServeTemplates bool
	ServeStatic    bool
}

type App struct {
	R      *gin.Engine
	Repos  *entities.Repositories
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
		err := app.runMigrations()

		if err != nil {
			return nil, err
		}
	}

	app.Repos, err = entities.NewRepositories(app.DBConn)
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
	controllers.SetUpRoutes(app.R)

	return app, nil
}
