package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.server.Close()

	tests := []struct {
		name    string
		urlPath string
		method  string
		body    string
		code    int
	}{
		{
			name:    "Create notification",
			urlPath: "/internal/v1/notifications",
			method:  http.MethodPost,
			body:    "{\"fromUserId\":1, \"toUserId\":2,\"postId\":1,\"type\":\"Like\"}",
			code:    201,
		},
		{
			name:    "Create notification with invalid type",
			urlPath: "/internal/v1/notifications",
			method:  http.MethodPost,
			body:    "{\"fromUserId\":1, \"toUserId\":2,\"postId\":1,\"type\":\"abc\"}",
			code:    400,
		},
		{
			name:    "Create notification with missing target user",
			urlPath: "/internal/v1/notifications",
			method:  http.MethodPost,
			body:    "{\"fromUserId\":1,\"postId\":1,\"type\":\"Like\"}",
			code:    400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := ts.request(t, tt.method, tt.urlPath, bytes.NewBuffer([]byte(tt.body)))

			if code != tt.code {
				t.Errorf("Got %d code, expected: %d", code, tt.code)
			}
		})
	}
}
