package main

import (
	"io"
	"log/slog"
	"nearby/common/clients"
	"nearby/common/httperrors"
	"nearby/common/middleware"
	"nearby/notifications/internal/data"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	logger := *slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{}))
	httpErrors := httperrors.NewHttpErrors(&logger)

	return &application{
		config:           config{},
		logger:           logger,
		models:           data.NewMockModels(),
		httpErrors:       httpErrors,
		commonMiddleware: middleware.NewCommonMiddleware(httpErrors),
		usersClient:      clients.MockUsersClient{},
	}
}

type testServer struct {
	server *httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	server := httptest.NewTLSServer(h)
	return &testServer{server}
}

func (ts *testServer) request(t *testing.T, method string, urlPath string, body io.Reader) int {
	requestUrl, err := url.Parse(ts.server.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(method, requestUrl.String(), body)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.server.Client().Do(request)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode
}
