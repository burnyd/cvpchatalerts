package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Used as a HTTP router for everything.

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/alert", app.CreateAlert)
	return router
}
