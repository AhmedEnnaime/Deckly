package createPrompt

import (
	"context"
	"deckly/cmd/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func validate(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var prompt models.Prompt
		if err := json.NewDecoder(r.Body).Decode(&prompt); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid JSON body")
			return
		}
		if len(strings.TrimSpace(prompt.Subject)) == 0 {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "Subject is required and cannot be empty")
			return
		}
		ctx := context.WithValue(r.Context(), models.CtxKey("newPrompt"), prompt)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
