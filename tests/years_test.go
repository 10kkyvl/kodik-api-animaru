package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestYears(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 30590,
  "results": [
    {
      "year": 2020,
      "count": 620
    },
    {
      "year": 2019,
      "count": 112
    }
  ]
}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Years(c, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 30590 {
		t.Errorf("expected Total=30590, got %d", result.Total)
	}
	if len(result.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(result.Results))
	}
	if result.Results[0].Year != 2020 {
		t.Errorf("expected first result 2020, got %d", result.Results[0].Year)
	}
}

func TestYears_WithParams(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "results": [
    {
      "year": 2020,
      "count": 620
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

	result, err := api.Years(c, &models.YearsParams{Types: "anime-serial"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected Total=1, got: %d", result.Total)
	}
}
