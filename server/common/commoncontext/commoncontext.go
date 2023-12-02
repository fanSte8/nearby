package commoncontext

import (
	"context"
	"net/http"
)

type contextKey string

const userIDContextKey = contextKey("userID")
const userActivatedContextKey = contextKey("userActivated")

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

func ContextSetUserActivated(r *http.Request, activated bool) *http.Request {
	ctx := context.WithValue(r.Context(), userActivatedContextKey, activated)
	return r.WithContext(ctx)
}

func ContextGetUserActivated(r *http.Request) bool {
	activated, ok := r.Context().Value(userActivatedContextKey).(bool)
	if !ok {
		panic("missing user activated value in request context")
	}

	return activated
}
