package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	r.Methods("POST").Path("/v1/mailer/activation").HandlerFunc(app.handleActivationTokenMail)
	r.Methods("POST").Path("/v1/mailer/password-reset").HandlerFunc(app.handlePasswordResetTokenMail)

	return app.commonMiddleware.RecoverPanic(r)
}
