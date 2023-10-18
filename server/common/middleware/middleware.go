package middleware

import (
	"fmt"
	"nearby/common/httperrors"
	"net/http"
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
