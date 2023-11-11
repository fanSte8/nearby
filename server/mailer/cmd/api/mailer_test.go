package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestEmails(t *testing.T) {
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
			name:    "Send mail with activation token",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\",\"recipient\":\"test@mail.com\"}",
			code:    200,
		},
		{
			name:    "Send mail with password reset token",
			urlPath: "/v1/mailer/password-reset",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\",\"recipient\":\"test@mail.com\"}",
			code:    200,
		},
		{
			name:    "Send mail with missing activation token",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"recipient\":\"test@mail.com\"}",
			code:    400,
		},

		{
			name:    "Send mail with missing password reset token",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"recipient\":\"test@mail.com\"}",
			code:    400,
		},

		{
			name:    "Send mail with missing activation recipient",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\"}",
			code:    400,
		},

		{
			name:    "Send mail with missing password reset recipient",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\"}",
			code:    400,
		},

		{
			name:    "Send mail with activation token with invalid recipient",
			urlPath: "/v1/mailer/activation",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\",\"recipient\":\"test\"}",
			code:    400,
		},
		{
			name:    "Send mail with password reset token  with invalid recipient",
			urlPath: "/v1/mailer/password-reset",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\",\"recipient\":\"test\"}",
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
