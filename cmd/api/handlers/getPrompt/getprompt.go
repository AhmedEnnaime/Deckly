package getPrompt

import (
	"database/sql"
	"deckly/cmd/api/models"
	"deckly/pkg/application"
	"deckly/pkg/logger"
	"deckly/pkg/middlewares"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getUser(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer r.Body.Close()
		id := r.Context().Value(models.CtxKey("promptid"))
		prompt := &models.Prompt{ID: id.(uuid.UUID)}
		if err := prompt.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "No such prompt found")
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(prompt)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	logger.Info.Printf("Sending request to get prompt by id")
	mdw := []middlewares.Middleware{
		middlewares.LogRequest,
		validateRequest,
	}
	return middlewares.Chain(getUser(app), mdw...)
}
