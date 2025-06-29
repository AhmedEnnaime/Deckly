package getPrompt

import (
	"deckly/pkg/application"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Do(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "Hello")
	}
}
