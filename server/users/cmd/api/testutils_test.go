package main

import (
	"io"
	"log/slog"
	"nearby/common/clients"
	"nearby/common/httperrors"
	"nearby/common/middleware"
	"nearby/common/storage"
	"nearby/users/internal/data"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	logger := *slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	httpErrors := httperrors.NewHttpErrors(&logger)

	return &application{
		config:           config{testing: true, JWTSecret: "test"},
		logger:           logger,
		models:           data.NewMockModels(),
		httpErrors:       httpErrors,
		commonMiddleware: middleware.NewCommonMiddleware(httpErrors),
		storage:          storage.MockStorage{},
		mailerClient:     clients.MockMailerClient{},
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
