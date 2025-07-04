package router

import (
	"deckly/cmd/api/handlers/getPrompt"
	"deckly/pkg/application"

	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/prompts/:id", getPrompt.Do(app))
	return mux
}
