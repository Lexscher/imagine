package createiuter

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/lexscher/imagine/cmd/api/models"
	"github.com/lexscher/imagine/pkg/application"
	"github.com/lexscher/imagine/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func createIuter(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		iuter := &models.Iuter{}
		json.NewDecoder(r.Body).Decode(iuter)

		if err := iuter.Create(r.Context(), app);err != nil {
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
	return middleware.Chain(createIuter(app), middleware.LogRequest)
}