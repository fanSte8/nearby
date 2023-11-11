package main

import (
	"nearby/common/commoncontext"
	"net/http"
)

func (app *application) authorize(handler http.HandlerFunc) http.HandlerFunc {
	if app.config.testing {
		return func(w http.ResponseWriter, r *http.Request) {
			r = commoncontext.ContextSetUserID(r, 1)
			handler.ServeHTTP(w, r)
		}
	}

	return app.commonMiddleware.Authorize(app.config.JWTSecret, handler)
}
