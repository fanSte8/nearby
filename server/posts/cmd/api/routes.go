package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	r.Methods(http.MethodPost).Path("/v1/posts").HandlerFunc(app.authorize(app.handleCreatePost))
	r.Methods(http.MethodGet).Path("/v1/posts/latest").HandlerFunc(app.authorize(app.handleGetLatestPosts))

	return app.commonMiddleware.RecoverPanic(r)
}
