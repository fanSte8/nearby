package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/v1/notifications").HandlerFunc(app.authorize(app.handleGetNotifications))
	r.Methods(http.MethodGet).Path("/v1/notifications/seen").HandlerFunc(app.authorize(app.handleHasUnseenNotifications))
	r.Methods(http.MethodPost).Path("/internal/v1/notifications").HandlerFunc(app.handleCreateNotification)

	return app.commonMiddleware.RecoverPanic(r)
}
