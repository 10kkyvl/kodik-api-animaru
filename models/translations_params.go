package models

import (
	"strconv"
	"strings"
)

// TranslationsParams описывает параметры запроса для эндпоинта /translations/v2 API Kodik.
// Если значение поля является пустым, оно не включается в итоговый запрос.
type TranslationsParams struct {
	// Основные параметры
	Types           string `json:"types,omitempty"`            // Типы материалов, например: anime-serial, foreign-movie, etc.
	Year            string `json:"year,omitempty"`             // Фильтрация по году
	TranslationType string `json:"translation_type,omitempty"` // Тип перевода: voice или subtitles
	HasField        string `json:"has_field,omitempty"`        // Наличие определённых полей (например, kinopoisk_id, imdb_id и т.д.)
	Lgbt            *bool  `json:"lgbt,omitempty"`             // Фильтр по наличию LGBT контента: true/false
	Sort            string `json:"sort,omitempty"`             // Сортировка результатов: title или count

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Страны через запятую (например: USA,Russia)
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую (например: action,drama)
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность или диапазон длительности (в минутах)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска или диапазон
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг или диапазон
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori или диапазон
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList или диапазон

	// Персоналии
	Actors    string `json:"actors,omitempty"`    // Список актёров через запятую
	Directors string `json:"directors,omitempty"` // Список режиссёров через запятую
	Producers string `json:"producers,omitempty"` // Список продюсеров через запятую
	Writers   string `json:"writers,omitempty"`   // Список сценаристов через запятую
	Composers string `json:"composers,omitempty"` // Список композиторов через запятую
	Editors   string `json:"editors,omitempty"`   // Список монтажёров через запятую
	Designers string `json:"designers,omitempty"` // Список дизайнеров через запятую
	Operators string `json:"operators,omitempty"` // Список операторов через запятую

	// Другие параметры
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA (например, PG, PG-13, R)
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст (например, 16 или 12-16)
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме: tv, movie, ova, ona и т.д.
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме: anons, ongoing, released и т.д.
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Лицензионные правообладатели через запятую
}

// ToMap преобразует структуру TranslationsParams в карту параметров для HTTP-запроса.
func (tp *TranslationsParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"types":              tp.Types,
		"year":               tp.Year,
		"translation_type":   tp.TranslationType,
		"has_field":          tp.HasField,
		"sort":               tp.Sort,
		"countries":          tp.Countries,
		"genres":             tp.Genres,
		"anime_genres":       tp.AnimeGenres,
		"drama_genres":       tp.DramaGenres,
		"all_genres":         tp.AllGenres,
		"duration":           tp.Duration,
		"kinopoisk_rating":   tp.KinopoiskRating,
		"imdb_rating":        tp.ImdbRating,
		"shikimori_rating":   tp.ShikimoriRating,
		"mydramalist_rating": tp.MydramalistRating,
		"actors":             tp.Actors,
		"directors":          tp.Directors,
		"producers":          tp.Producers,
		"writers":            tp.Writers,
		"composers":          tp.Composers,
		"editors":            tp.Editors,
		"designers":          tp.Designers,
		"operators":          tp.Operators,
		"rating_mpaa":        tp.RatingMPAA,
		"minimal_age":        tp.MinimalAge,
		"anime_kind":         tp.AnimeKind,
		"mydramalist_tags":   tp.MydramalistTags,
		"anime_status":       tp.AnimeStatus,
		"drama_status":       tp.DramaStatus,
		"all_status":         tp.AllStatus,
		"anime_studios":      tp.AnimeStudios,
		"anime_licensed_by":  tp.AnimeLicensedBy,
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

	if tp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*tp.Lgbt)
	}

	return params
}
