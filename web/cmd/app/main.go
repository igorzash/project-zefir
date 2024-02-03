package main

import "github.com/igorzash/project-zefir/web/app"

func main() {
	app, err := app.NewApp(&app.AppParams{
		ServeStatic:    true,
		ServeTemplates: true,
	})
	if err != nil {
		panic(err)
	}

	app.R.Run("0.0.0.0:8080")
}
