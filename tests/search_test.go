package tests

import (
	"net/http"
	"testing"

	"github.com/10kkyvl/kodik-api-animaru/api"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

func TestSearch(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "results": [
    {
      "id": "movie-452654",
      "type": "foreign-movie",
      "link": "https://example.com/mock-link/720p",
      "title": "Аватар",
      "title_orig": "Avatar",
      "translation": {
        "id": 704,
        "title": "Дублированный",
        "type": "voice"
      },
      "year": 2009,
      "kinopoisk_id": "251733",
      "imdb_id": "tt0499549",
      "quality": "BDRip 720p",
      "blocked_countries": [],
      "created_at": "2014-06-22T22:19:22Z",
      "updated_at": "2016-04-25T07:03:33Z"
    }
  ]
}`
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Search(c, &models.SearchParams{Title: "Аватар"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(result.Results))
	}
	if result.Results[0].Title != "Аватар" {
		t.Errorf("expected result Аватар, got %s", result.Results[0].Title)
	}
}

func TestSearch_WithFullParams(t *testing.T) {
	fixture := `{
  "time": "5ms",
  "total": 1,
  "results": [
    {
      "id": "serial-452654",
      "type": "anime",
      "link": "https://example.com/serial/mock-uuid/720p",
      "title": "Игра престолов",
      "title_orig": "Game of Thrones",
      "translation": {
        "id": 611,
        "title": "ColdFilm",
        "type": "voice"
      },
      "year": 2011,
      "last_season":9,
      "last_episode":4,
      "episodes_count":119,
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
		if r.URL.Query().Get("title") != "Игра престолов" {
			t.Errorf("expected title=Игра престолов, got %q", r.URL.Query().Get("title"))
		}
		if r.URL.Query().Get("strict") != "true" {
			t.Errorf("expected strict=true, got %q", r.URL.Query().Get("strict"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fixture))
	})

	result, err := api.Search(c, &models.SearchParams{Title: "Игра престолов", Strict: true})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("expected 1 result, got: %d", len(result.Results))
	}
}
