package main

import (
	"bytes"
	"fmt"
	"nearby/users/internal/data"
	"net/http"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.server.Close()

	tests := []struct {
		name    string
		urlPath string
		method  string
		code    int
	}{
		{
			name:    "Get existing user",
			urlPath: "/v1/users/1",
			method:  http.MethodGet,
			code:    200,
		},
		{
			name:    "User not found",
			urlPath: "/v1/users/2",
			method:  http.MethodGet,
			code:    404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := ts.request(t, tt.method, tt.urlPath, nil)

			if code != tt.code {
				t.Errorf("Got %d code, expected: %d", code, tt.code)
			}
		})
	}
}

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
			name:    "New user",
			urlPath: "/v1/users/register",
			method:  http.MethodPost,
			body:    "{\"firstName\":\"Test\",\"lastName\":\"Test\",\"email\":\"new@mail.com\",\"password\":\"1234567890\"}",
			code:    201,
		},
		{
			name:    "User with existing email",
			urlPath: "/v1/users/register",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"firstName\":\"Test\",\"lastName\":\"Test\",\"email\":\"%s\",\"password\":\"1234567890\"}", data.MockUserEmail),
			code:    400,
		},
		{
			name:    "Invalid email",
			urlPath: "/v1/users/register",
			method:  http.MethodPost,
			body:    "{\"firstName\":\"Test\",\"lastName\":\"Test\",\"email\":\"mail\",\"password\":\"1234567890\"}",
			code:    400,
		},
		{
			name:    "Short password",
			urlPath: "/v1/users/register",
			method:  http.MethodPost,
			body:    "{\"firstName\":\"Test\",\"lastName\":\"Test\",\"email\":\"new@mail.com\",\"password\":\"123\"}",
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

func TestLoginUser(t *testing.T) {
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
			name:    "Login",
			urlPath: "/v1/users/login",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"email\":\"%s\",\"password\":\"%s\"}", data.MockUserEmail, data.MockUserPassword),
			code:    200,
		},
		{
			name:    "Login with incorrect email",
			urlPath: "/v1/users/login",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"email\":\"incorrect@mail.com\",\"password\":\"%s\"}", data.MockUserPassword),
			code:    401,
		},
		{
			name:    "Login with incorrect password",
			urlPath: "/v1/users/login",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"email\":\"%s\",\"password\":\"12341234\"}", data.MockUserEmail),
			code:    401,
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

func TestChangePassword(t *testing.T) {
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
			name:    "Change password",
			urlPath: "/v1/users/change-password",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"oldPassword\":\"%s\",\"password\":\"12341234\"}", data.MockUserPassword),
			code:    200,
		},
		{
			name:    "Change password with incorrect old password",
			urlPath: "/v1/users/change-password",
			method:  http.MethodPost,
			body:    "{\"oldPassword\":\"1111111111\",\"password\":\"12341234\"}",
			code:    400,
		},
		{
			name:    "Change password with invalid new password",
			urlPath: "/v1/users/change-password",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"oldPassword\":\"%s\",\"password\":\"1111\"}", data.MockUserPassword),
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

func TestResetPassword(t *testing.T) {
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
			name:    "Reset password",
			urlPath: "/v1/users/reset-password",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"code\":\"%s\",\"password\":\"12341234\"}", data.MockPasswordResetToken),
			code:    200,
		},
		{
			name:    "Reset password with incorrect code",
			urlPath: "/v1/users/reset-password",
			method:  http.MethodPost,
			body:    "{\"code\":\"abc\",\"password\":\"12341234\"}",
			code:    400,
		},
		{
			name:    "Reset password with invalid new password",
			urlPath: "/v1/users/reset-password",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"code\":\"%s\",\"password\":\"1111\"}", data.MockPasswordResetToken),
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

func TestActivationToken(t *testing.T) {
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
			name:    "Activate user",
			urlPath: "/v1/users/activate",
			method:  http.MethodPost,
			body:    fmt.Sprintf("{\"token\":\"%s\"}", data.MockActivationToken),
			code:    200,
		},
		{
			name:    "Activate user with incorrect token",
			urlPath: "/v1/users/activate",
			method:  http.MethodPost,
			body:    "{\"token\":\"abc\"}",
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
