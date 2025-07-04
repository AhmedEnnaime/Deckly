package getPrompt

import (
	"context"
	"deckly/cmd/api/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func validateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uidStr := p.ByName("id")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "malformed id")
			return
		}
		ctx := context.WithValue(r.Context(), models.CtxKey("promptid"), uid)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
