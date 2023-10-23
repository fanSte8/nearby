package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	router.HandlerFunc(http.MethodPost, "/v1/mailer/activation", app.handleActivationToken)
	router.HandlerFunc(http.MethodPost, "/v1/mailer/password-reset", app.handlePasswordResetToken)

	return app.commonMiddleware.RecoverPanic(router)
}
