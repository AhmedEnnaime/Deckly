package api

import (
	"deckly/pkg/application"
	"deckly/pkg/exitHandler"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env vars")
	}
	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}
	exitHandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
