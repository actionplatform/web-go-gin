package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func call(t *testing.T, r http.Handler, method, path, body string) (int, map[string]any) {
	t.Helper()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	var out map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &out); err != nil {
		out = map[string]any{"raw": w.Body.String()}
	}
	return w.Code, out
}

func TestPing(t *testing.T) {
	code, body := call(t, New(), http.MethodGet, "/health", "")
	if code != http.StatusOK || body["status"] != "ok" || body["version"] != Version {
		t.Fatalf("got %d %v", code, body)
	}
}

func TestCreateAndGetItem(t *testing.T) {
	r := New()
	code, created := call(t, r, http.MethodPost, "/api/v1/items", `{"name": "pen"}`)
	if code != http.StatusCreated || created["name"] != "pen" {
		t.Fatalf("got %d %v", code, created)
	}
	code, got := call(t, r, http.MethodGet, "/api/v1/items/1", "")
	if code != http.StatusOK || got["name"] != "pen" {
		t.Fatalf("got %d %v", code, got)
	}
}

func TestErrors(t *testing.T) {
	r := New()
	if code, _ := call(t, r, http.MethodPost, "/api/v1/items", `{"name": "  "}`); code != http.StatusUnprocessableEntity {
		t.Fatalf("invalid name: got %d", code)
	}
	if code, _ := call(t, r, http.MethodGet, "/api/v1/items/999", ""); code != http.StatusNotFound {
		t.Fatalf("unknown id: got %d", code)
	}
}
