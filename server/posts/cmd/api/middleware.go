package main

import "net/http"

func (app *application) authorize(handler http.HandlerFunc) http.HandlerFunc {
	return app.commonMiddleware.Authorize(app.config.JWTSecret, handler)
}
