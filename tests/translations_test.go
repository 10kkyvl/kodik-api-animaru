package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestTranslations(t *testing.T) {
	fixture := `{"time":"1ms","total":2,"results":[{"id":1, "title": "SoftBox", "count": 1000},{"id":2, "title": "AniLibria", "count": 500}]}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Translations(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("expected Total=2, got %d", result.Total)
	}
	if len(result.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(result.Results))
	}
	if result.Results[0].Title != "SoftBox" {
		t.Errorf("expected first result SoftBox, got %s", result.Results[0].Title)
	}
}

func TestTranslations_WithParams(t *testing.T) {
	fixture := `{"time":"1ms","total":1,"results":[{"id":1,"title":"SoftBox","count":1000}]}`

	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("types") != "anime-serial" {
			t.Errorf("expected types=anime-serial, got %q", r.URL.Query().Get("types"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Translations(c, &models.TranslationsParams{Types: "anime-serial"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
