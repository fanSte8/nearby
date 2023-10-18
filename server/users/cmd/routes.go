package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	router.HandlerFunc(http.MethodPost, "/", app.handler)

	return app.commonMiddleware.RecoverPanic(router)
}
