package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"afonso.dev/snippetbox/internal/assert"
)

func TestPingV1(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}

func TestPingV2(t *testing.T) {
	app := &application{
		logger: slog.New(slog.DiscardHandler),
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	res := ts.get(t, "/ping")
	assert.Equal(t, res.status, http.StatusOK)
	assert.Equal(t, res.body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name       string
		urlPath    string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Valid ID",
			urlPath:    "/snippet/view/1",
			wantStatus: http.StatusOK,
			wantBody:   "An old silent pond...",
		},
		{
			name:       "Non-existent ID",
			urlPath:    "/snippet/view/2",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Negative ID",
			urlPath:    "/snippet/view/-1",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Decimal ID",
			urlPath:    "/snippet/view/1.23",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "String ID",
			urlPath:    "/snippet/view/foo",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Empty ID",
			urlPath:    "/snippet/view/",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts.resetClientCookieJar(t)

			res := ts.get(t, tt.urlPath)
			assert.Equal(t, res.status, tt.wantStatus)
			assert.True(t, strings.Contains(res.body, tt.wantBody))
		})
	}
}
