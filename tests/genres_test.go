package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestGenres(t *testing.T) {
	fixture := `{"time":"1ms","total":2,"results":[{"title":"Action", "count": 100},{"title":"Comedy", "count": 80}]}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Genres(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("expected Total=2, got %d", result.Total)
	}
	if len(result.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(result.Results))
	}
	if result.Results[0].Title != "Action" {
		t.Errorf("expected first result Action, got %s", result.Results[0].Title)
	}
}

func TestGenres_WithParams(t *testing.T) {
	fixture := `{"time":"1ms","total":1,"results":[{"title":"Action","count":100}]}`

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("types") != "anime" {
			t.Errorf("expected types=anime, got %q", r.URL.Query().Get("types"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Genres(c, &models.GenresParams{Types: "anime"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
