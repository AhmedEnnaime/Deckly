package createPrompt

import (
	"deckly/cmd/api/models"
	"deckly/pkg/application"
	"deckly/pkg/middlewares"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreatePrompt(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()
		val := r.Context().Value(models.CtxKey("newPrompt"))
		prompt, ok := val.(models.Prompt)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid prompt data")
			return
		}
		if err := prompt.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to create prompt")
			return
		}
		go func(p models.Prompt) {
			if err := p.TriggerN8nWorkflow(app); err != nil {
				fmt.Println("Failed to trigger n8n:", err)
			}
		}(prompt)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(prompt)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middlewares.Middleware{
		middlewares.LogRequest,
		validate,
	}
	return middlewares.Chain(CreatePrompt(app), mdw...)
}
