package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	r.Methods(http.MethodPost).Path("/v1/posts").HandlerFunc(app.authorize(app.handleCreatePost))
	r.Methods(http.MethodDelete).Path("/v1/posts/{id}").HandlerFunc(app.authorize(app.handleDeletePost))
	r.Methods(http.MethodGet).Path("/v1/posts/latest").HandlerFunc(app.authorize(app.handleGetLatestPosts))

	r.Methods(http.MethodPost).Path("/v1/posts/{postId}/likes").HandlerFunc(app.authorize(app.handlePostLike))

	r.Methods(http.MethodGet).Path("/v1/posts/{postId}/comments").HandlerFunc(app.authorize(app.handleGetComments))
	r.Methods(http.MethodPost).Path("/v1/posts/{postId}/comments").HandlerFunc(app.authorize(app.handleCreateComment))
	r.Methods(http.MethodDelete).Path("/v1/posts/comments/{id}").HandlerFunc(app.authorize(app.handleDeleteComment))

	return app.commonMiddleware.RecoverPanic(r)
}
