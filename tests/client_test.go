package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/client"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) *client.Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	c := client.NewClient("test-token", false)
	c.BaseURL = server.URL
	return c
}

func TestDoRequest_GET_Success(t *testing.T) {
	type Response struct {
		Total int `json:"total"`
	}

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") != "test-token" {
			t.Errorf("expected token=test-token, got %q", r.URL.Query().Get("token"))
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Total: 42})
	})

	var result Response
	err := c.DoRequest("GET", "/countries", nil, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 42 {
		t.Errorf("expected Total=42, got %d", result.Total)
	}
}

func TestDoRequest_NonOKStatus(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})

	var result struct{}
	err := c.DoRequest("GET", "/countries", nil, &result)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDoRequest_UnsupportedMethod(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {})

	var result struct{}
	err := c.DoRequest("DELETE", "/countries", nil, &result)

	if err == nil {
		t.Fatal("expected error for unsupported method")
	}
}
