package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestQualities(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 2,
  "results": [
    {
      "title": "WEB-DLRip 720p",
      "count": 545
    },
    {
      "title": "WEB-DLRip",
      "count": 11
    }
  ]
}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Qualities(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("expected Total=2, got %d", result.Total)
	}
	if len(result.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(result.Results))
	}
	if result.Results[0].Title != "WEB-DLRip 720p" {
		t.Errorf("expected first result WEB-DLRip 720p, got %s", result.Results[0].Title)
	}
}

func TestQualities_WithParams(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "results": [
    {
      "title": "WEB-DLRip 720p",
      "count": 545
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

	result, err := api.Qualities(c, &models.QualitiesParams{Types: "anime-serial"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
