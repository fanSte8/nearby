package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestCreatePost(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.server.Close()

	file, err := os.Open("../../internal/test_image.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	t.Run("Create new post", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		_ = writer.WriteField("description", "Post")
		_ = writer.WriteField("longitude", "10")
		_ = writer.WriteField("latitude", "10")

		part, err := writer.CreateFormFile("image", "image.png")
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		code := ts.request(t, http.MethodPost, "/v1/posts", body, map[string]string{"Content-Type": writer.FormDataContentType()})

		if code != 201 {
			t.Errorf("Got %d code, expected: %d", code, 201)
		}
	})

	t.Run("Create new post with invalid latitude", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		_ = writer.WriteField("description", "Post")
		_ = writer.WriteField("longitude", "10")
		_ = writer.WriteField("latitude", "abc")

		part, err := writer.CreateFormFile("image", "image.png")
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		code := ts.request(t, http.MethodPost, "/v1/posts", body, map[string]string{"Content-Type": writer.FormDataContentType()})

		if code != 400 {
			t.Errorf("Got %d code, expected: %d", code, 400)
		}
	})

	t.Run("Create new post with invalid longitude", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		_ = writer.WriteField("description", "Post")
		_ = writer.WriteField("longitude", "abc")
		_ = writer.WriteField("latitude", "10")

		part, err := writer.CreateFormFile("image", "image.png")
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		code := ts.request(t, http.MethodPost, "/v1/posts", body, map[string]string{"Content-Type": writer.FormDataContentType()})

		if code != 400 {
			t.Errorf("Got %d code, expected: %d", code, 400)
		}
	})

	t.Run("Create new post with empty description", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		_ = writer.WriteField("description", "Post")
		_ = writer.WriteField("longitude", "abc")
		_ = writer.WriteField("latitude", "10")

		part, err := writer.CreateFormFile("image", "image.png")
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		code := ts.request(t, http.MethodPost, "/v1/posts", body, map[string]string{"Content-Type": writer.FormDataContentType()})

		if code != 400 {
			t.Errorf("Got %d code, expected: %d", code, 400)
		}
	})
}

func TestGetPostByID(t *testing.T) {
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
			name:    "Get post by id",
			urlPath: "/v1/posts/1?longitude=10&latitude=10",
			method:  http.MethodGet,
			code:    200,
		},
		{
			name:    "Post not found",
			urlPath: "/v1/posts/123?longitude=10&latitude=10",
			method:  http.MethodGet,
			code:    404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := ts.request(t, tt.method, tt.urlPath, nil, nil)

			if code != tt.code {
				t.Errorf("Got %d code, expected: %d", code, tt.code)
			}
		})
	}
}

func TestLikePost(t *testing.T) {
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
			name:    "Like post",
			urlPath: "/v1/posts/1/likes",
			method:  http.MethodPost,
			code:    200,
		},
		{
			name:    "Like non-existing post",
			urlPath: "/v1/posts/123/likes",
			method:  http.MethodPost,
			code:    404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := ts.request(t, tt.method, tt.urlPath, nil, nil)

			if code != tt.code {
				t.Errorf("Got %d code, expected: %d", code, tt.code)
			}
		})
	}
}

func TestCreateComment(t *testing.T) {
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
			name:    "Create comment",
			urlPath: "/v1/posts/1/comments",
			method:  http.MethodPost,
			body:    "{\"text\":\"comment\"}",
			code:    200,
		},
		{
			name:    "Create comment with empty text",
			urlPath: "/v1/posts/1/comments",
			method:  http.MethodPost,
			body:    "{\"text\":\"\"}",
			code:    400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := ts.request(t, tt.method, tt.urlPath, bytes.NewBuffer([]byte(tt.body)), nil)

			if code != tt.code {
				t.Errorf("Got %d code, expected: %d", code, tt.code)
			}
		})
	}
}
