package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.loginHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/change-password", app.authorize(app.handleChangePassword))

	return app.commonMiddleware.RecoverPanic(router)
}
