package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandle ensures that Handle executes without error and returns the
// HTTP 200 status code indicating no errors.
func TestHandle(t *testing.T) {
	// initialize request body
	var body io.Reader = io.NopCloser(strings.NewReader(`{"x": 12, "y": 23.3}`))

	var (
		w   = httptest.NewRecorder()
		req = httptest.NewRequest("POST", "http://example.com/test", body)
		res *http.Response
	)

	handler(w, req)
	res = w.Result()
	var result ResponseStruct
	json.NewDecoder(res.Body).Decode(&result)

	if result.Result != 23.3 {
		t.Fatalf("unexpected result: %v", result.Result)
	}
}
