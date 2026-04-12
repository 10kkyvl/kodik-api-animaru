package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestList(t *testing.T) {
	fixture := `{"time":"1ms","total":1,"results":[{"title":"The Matrix", "id": "movie-2"}]}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.List(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("expected Total=1, got %d", result.Total)
	}
}

func TestList_WithParams(t *testing.T) {
	fixture := `{"time":"1ms","total":1,"results":[{"title":"The Matrix","id":"movie-2"}]}`

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("types") != "foreign-movie" {
			t.Errorf("expected types=foreign-movie, got %q", r.URL.Query().Get("types"))
		}
		if r.URL.Query().Get("limit") != "1" {
			t.Errorf("expected limit=1, got %q", r.URL.Query().Get("limit"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.List(c, &models.ListParams{Types: "foreign-movie", Limit: 1})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
