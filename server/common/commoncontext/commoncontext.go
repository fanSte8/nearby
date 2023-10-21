package commoncontext

import (
	"context"
	"net/http"
)

type contextKey string

const userIDContextKey = contextKey("userID")

func ContextSetUserID(r *http.Request, userID int64) *http.Request {
	ctx := context.WithValue(r.Context(), userIDContextKey, userID)
	return r.WithContext(ctx)
}

func ContextGetUserID(r *http.Request) int64 {
	userID, ok := r.Context().Value(userIDContextKey).(int64)
	if !ok {
		panic("missing user value in request context")
	}

	return userID
}
