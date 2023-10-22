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
	router.HandlerFunc(http.MethodPost, "/v1/users/forgotten-password", app.handleForgottenPassword)
	router.HandlerFunc(http.MethodPost, "/v1/users/reset-password", app.handleResetPassword)

	router.HandlerFunc(http.MethodGet, "/v1/users/activate", app.authorize(app.handleNewActivationToken))
	router.HandlerFunc(http.MethodPost, "/v1/users/activate", app.authorize(app.handleActivateAccount))

	return app.commonMiddleware.RecoverPanic(router)
}
