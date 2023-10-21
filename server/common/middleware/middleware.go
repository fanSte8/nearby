package middleware

import (
	"fmt"
	"nearby/common/commoncontext"
	"nearby/common/httperrors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

type CommonMiddleware struct {
	errors httperrors.HttpErrors
}

func NewCommonMiddleware(errors httperrors.HttpErrors) CommonMiddleware {
	return CommonMiddleware{errors}
}

func (m CommonMiddleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				m.errors.ServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (m CommonMiddleware) Authorize(JWTSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")

		if authorizationHeader == "" {
			m.errors.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			m.errors.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(JWTSecret))
		if err != nil {
			m.errors.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		if !claims.Valid(time.Now()) {
			m.errors.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			m.errors.ServerErrorResponse(w, r, err)
			return
		}

		r = commoncontext.ContextSetUserID(r, userID)

		next.ServeHTTP(w, r)
	})
}
