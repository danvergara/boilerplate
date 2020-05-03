package main

import (
	"github.com/danvergara/boilerplate/cmd/api/router"
	"github.com/danvergara/boilerplate/pkg/application"
	"github.com/danvergara/boilerplate/pkg/exithandler"
	"github.com/danvergara/boilerplate/pkg/logger"
	"github.com/danvergara/boilerplate/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Info.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			logger.Error.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
