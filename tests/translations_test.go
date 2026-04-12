package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestTranslations(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 2,
  "results": [
    {
      "id": 735,
      "title": "2x2",
      "count": 26
    },
    {
      "id": 824,
      "title": "3df voice",
      "count": 16
    }
  ]
}`
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
	if result.Results[0].Title != "2x2" {
		t.Errorf("expected first result 2x2, got %s", result.Results[0].Title)
	}
}

func TestTranslations_WithParams(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "results": [
    {
      "id": 735,
      "title": "2x2",
      "count": 26
    }
  ]
}`

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
