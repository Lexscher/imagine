package getiuter

import (
	"fmt"
	"errors"
	"net/http"
	"database/sql"
	"encoding/json"

	"github.com/lexscher/imagine/cmd/api/models"
	"github.com/lexscher/imagine/pkg/application"
	"github.com/lexscher/imagine/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func getIuter(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("iuterid"))
		iuter := &models.Iuter{ID: id.(int)}

		if err := iuter.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "that iuter does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Oopsie!")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(iuter)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getIuter(app), mdw...)
}