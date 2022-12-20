package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"github.com/home-first/pocketKit/frontend"
)

func main() {
	stripped, err := fs.Sub(frontend.Assets, "dist")
	if err != nil {
		log.Fatal(err)
		return
	}

	frontendHandler := http.FileServer(http.FS(stripped))
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", echo.WrapHandler(frontendHandler))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
