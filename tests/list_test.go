package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestList(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "prev_page": null,
  "next_page": null,
  "results": [
    {
      "id": "serial-452654",
      "type": "anime",
      "link": "http://kodikplayer.com/serial/4309/bc6def495a31545c7f648f7fb68d22a8/720p",
      "title": "Игра престолов",
      "title_orig": "Game of Thrones",
      "translation": {
        "id": 611,
        "title": "ColdFilm",
        "type": "voice"
      },
      "year": 2011,
      "last_season": 9,
      "last_episode": 4,
      "episodes_count": 119,
      "kinopoisk_id": "1161904",
      "imdb_id": "tt0944947",
      "quality": "WEB-DL 720p",
      "blocked_countries": ["RU"],
      "created_at": "2017-07-17T16:34:52Z",
      "updated_at": "2018-04-06T14:19:27Z"
    }
  ]
}`
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
	if result.Results[0].Title != "Игра престолов" {
		t.Errorf("expected title Игра престолов, got %s", result.Results[0].Title)
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
