package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	router.HandlerFunc(http.MethodPost, "/v1/user/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/user/login", app.loginHandler)

	return app.commonMiddleware.RecoverPanic(router)
}
