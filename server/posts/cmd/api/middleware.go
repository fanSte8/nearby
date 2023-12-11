package main

import (
	"nearby/common/commoncontext"
	"net/http"
)

func (app *application) authorize(handler http.HandlerFunc) http.HandlerFunc {
	if app.config.testing {
		return func(w http.ResponseWriter, r *http.Request) {
			r = commoncontext.ContextSetUserID(r, 1)
			r = commoncontext.ContextSetUserActivated(r, true)
			handler.ServeHTTP(w, r)
		}
	}

	return app.commonMiddleware.Authorize(app.config.JWTSecret, handler)
}

func (app *application) isActivated(handler http.HandlerFunc) http.HandlerFunc {
	return app.commonMiddleware.IsActivated(handler)
}
