package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestCountries(t *testing.T) {
	fixture := `{"time":"1ms","total":3,"results":[{"title":"Japan", "count": 620},{"title":"USA", "count": 135},{"title":"Korea", "count": 120}]}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Countries(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 3 {
		t.Errorf("expected Total=3, got %d", result.Total)
	}
	if len(result.Results) != 3 {
		t.Errorf("expected 3 results, got %d", len(result.Results))
	}
	if result.Results[0].Title != "Japan" {
		t.Errorf("expected first result Japan, got %s", result.Results[0].Title)
	}
	_ = c
}

func TestCountries_WithParams(t *testing.T) {
	fixture := `{"time":"1ms","total":1,"results":[{"title":"Japan","count":620}]}`

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("types") != "anime" {
			t.Errorf("expected types=anime, got %q", r.URL.Query().Get("types"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Countries(c, &models.CountriesParams{Types: "anime"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
