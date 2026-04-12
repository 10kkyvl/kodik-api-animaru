package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestSearch(t *testing.T) {
	fixture := `{"time":"1ms","results":[{"title":"Terminator", "id": "movie-1"}]}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Search(c, &models.SearchParams{Title: "Terminator"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(result.Results))
	}
	if result.Results[0].Title != "Terminator" {
		t.Errorf("expected result Terminator, got %s", result.Results[0].Title)
	}
}

func TestSearch_WithFullParams(t *testing.T) {
	fixture := `{"time":"1ms","results":[{"title":"Terminator","id":"movie-1"}]}`

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("title") != "Terminator" {
			t.Errorf("expected title=Terminator, got %q", r.URL.Query().Get("title"))
		}
		if r.URL.Query().Get("strict") != "true" {
			t.Errorf("expected strict=true, got %q", r.URL.Query().Get("strict"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Search(c, &models.SearchParams{Title: "Terminator", Strict: true})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("expected 1 result, got: %d", len(result.Results))
	}
}
