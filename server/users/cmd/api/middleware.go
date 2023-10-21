package main

import (
	"errors"
	"nearby/common/commoncontext"
	"nearby/users/internal/data"
	"net/http"
)

func (app *application) authorize(next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := commoncontext.ContextGetUserID(r)

		user, err := app.models.Users.GetById(userID)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.httpErrors.InvalidAuthenticationTokenResponse(w, r)
			default:
				app.httpErrors.ServerErrorResponse(w, r, err)
			}
			return
		}

		r = app.contextSetUser(r, user)

		next.ServeHTTP(w, r)
	})

	return app.commonMiddleware.Authorize(app.config.JWTSecret, fn)
}
