package models

import (
	"strconv"
	"strings"
)

// QualitiesParams описывает параметры запроса для эндпоинта /qualities/v2 API Kodik.
// Если значение поля является пустым или нулевым, оно не включается в итоговый запрос.
type QualitiesParams struct {
	// Основные параметры
	Types           string `json:"types,omitempty"`            // Типы материалов (например, "anime-serial,foreign-movie,..." )
	Year            string `json:"year,omitempty"`             // Год или диапазон (например, "2020", "2010-2020")
	TranslationID   int    `json:"translation_id,omitempty"`   // ID озвучки
	TranslationType string `json:"translation_type,omitempty"` // Тип перевода: voice или subtitles
	HasField        string `json:"has_field,omitempty"`        // Наличие определённых полей (например, kinopoisk_id)
	Lgbt            *bool  `json:"lgbt,omitempty"`             // Фильтр по LGBT контенту (true/false)
	Sort            string `json:"sort,omitempty"`             // Сортировка результатов: title или count

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Страны через запятую (например, "USA,Russia")
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность или диапазон (например, "90", "80-120")
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (например, "7.0", "6.5-8.2")
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList

	// Персоналии (список имён через запятую)
	Actors    string `json:"actors,omitempty"`
	Directors string `json:"directors,omitempty"`
	Producers string `json:"producers,omitempty"`
	Writers   string `json:"writers,omitempty"`
	Composers string `json:"composers,omitempty"`
	Editors   string `json:"editors,omitempty"`
	Designers string `json:"designers,omitempty"`
	Operators string `json:"operators,omitempty"`

	// Другие параметры
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA (например, "PG-13")
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст (например, "16", "12-16")
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме (tv, movie, ova, ona и т.д.)
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме (anons, ongoing, released, etc.)
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Лицензионные правообладатели через запятую
}

// ToMap преобразует структуру QualitiesParams в карту параметров для HTTP-запроса.
func (qp *QualitiesParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"types":              qp.Types,
		"year":               qp.Year,
		"translation_type":   qp.TranslationType,
		"has_field":          qp.HasField,
		"sort":               qp.Sort,
		"countries":          qp.Countries,
		"genres":             qp.Genres,
		"anime_genres":       qp.AnimeGenres,
		"drama_genres":       qp.DramaGenres,
		"all_genres":         qp.AllGenres,
		"duration":           qp.Duration,
		"kinopoisk_rating":   qp.KinopoiskRating,
		"imdb_rating":        qp.ImdbRating,
		"shikimori_rating":   qp.ShikimoriRating,
		"mydramalist_rating": qp.MydramalistRating,
		"actors":             qp.Actors,
		"directors":          qp.Directors,
		"producers":          qp.Producers,
		"writers":            qp.Writers,
		"composers":          qp.Composers,
		"editors":            qp.Editors,
		"designers":          qp.Designers,
		"operators":          qp.Operators,
		"rating_mpaa":        qp.RatingMPAA,
		"minimal_age":        qp.MinimalAge,
		"anime_kind":         qp.AnimeKind,
		"mydramalist_tags":   qp.MydramalistTags,
		"anime_status":       qp.AnimeStatus,
		"drama_status":       qp.DramaStatus,
		"all_status":         qp.AllStatus,
		"anime_studios":      qp.AnimeStudios,
		"anime_licensed_by":  qp.AnimeLicensedBy,
	}

	for k, v := range fields {
		if v != "" {
			if k == "countries" {
				params[k] = strings.ReplaceAll(v, " ", "")
			} else {
				params[k] = v
			}
		}
	}

	if qp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(qp.TranslationID)
	}
	if qp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*qp.Lgbt)
	}

	return params
}
