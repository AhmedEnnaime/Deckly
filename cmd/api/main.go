package main

import (
	"deckly/cmd/router"
	"deckly/pkg/application"
	"deckly/pkg/exitHandler"
	"deckly/pkg/logger"
	"deckly/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Info.Println("Failed to load env vars")
	}
	app, err := application.Get()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}
	srv := server.Get().WithAddr(app.Cfg.GetAPIPort()).WithRouter(router.Get(app)).WithErrLogger(logger.Error)
	go func() {
		logger.Info.Printf("Starting Server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()
	exitHandler.Init(func() {
		if err := srv.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
		if err := app.DB.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
