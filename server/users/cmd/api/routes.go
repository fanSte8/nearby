package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.httpErrors.NotFoundResponse)

	r.Methods(http.MethodPost).Path("/v1/users/register").HandlerFunc(app.handleRegisterUser)
	r.Methods(http.MethodPost).Path("/v1/users/login").HandlerFunc(app.handleLogin)

	r.Methods(http.MethodGet).Path("/v1/users/me").HandlerFunc(app.authorize(app.handleCurrentUserData))

	r.Methods(http.MethodPost).Path("/v1/users/change-password").HandlerFunc(app.authorize(app.handleChangePassword))
	r.Methods(http.MethodPost).Path("/v1/users/forgotten-password").HandlerFunc(app.handleForgottenPassword)
	r.Methods(http.MethodPost).Path("/v1/users/reset-password").HandlerFunc(app.authorize(app.handleResetPassword))

	r.Methods(http.MethodGet).Path("/v1/users/activate").HandlerFunc(app.authorize(app.handleNewActivationToken))
	r.Methods(http.MethodPost).Path("/v1/users/activate").HandlerFunc(app.authorize(app.handleActivateAccount))

	r.Methods(http.MethodPost).Path("/v1/users/profile-picture").HandlerFunc(app.authorize(app.handleProfilePictureUpload))

	r.Methods(http.MethodGet).Path("/internal/v1/users/{id}").HandlerFunc(app.handleGetUserByID)
	r.Methods(http.MethodGet).Path("/v1/users/{id}").HandlerFunc(app.authorize(app.handleGetUserByID))

	r.Methods(http.MethodPost).Path("/v1/users/posts-radius").HandlerFunc(app.authorize(app.handlePostsRadiusUpdate))

	return app.commonMiddleware.RecoverPanic(r)
}
