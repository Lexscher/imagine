package getiuter

import (
	"fmt"
	"context"
	"strconv"
	"net/http"

	"github.com/lexscher/imagine/cmd/api/models"
	"github.com/julienschmidt/httprouter"
)

func validateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		iuid := p.ByName("id")

		id, err := strconv.Atoi(iuid)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("iuterid"), id)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}