package router

import (
	"github.com/lexscher/imagine/cmd/api/handlers/createiuter"
	"github.com/lexscher/imagine/cmd/api/handlers/getiuter"
	"github.com/lexscher/imagine/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/iuters/:id", getiuter.Do(app))
	mux.POST("/iuters", createiuter.Do(app))
	return mux
}