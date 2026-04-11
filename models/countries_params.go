package models

import (
	"strconv"
	"strings"
)

// CountriesParams описывает параметры запроса для эндпоинта /countries API Kodik.
type CountriesParams struct {
	// Фильтрация материалов
	Types             string `json:"types,omitempty"`              // Тип материала (например, "foreign-movie,cartoon-serial")
	Year              string `json:"year,omitempty"`               // Год или диапазон (например, "2020" или "2010-2020")
	TranslationID     int    `json:"translation_id,omitempty"`     // ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // IDs озвучек, которые нужно исключить (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода (voice или subtitles)
	HasField          string `json:"has_field,omitempty"`          // Наличие определённых полей (например: kinopoisk_id, imdb_id, mdl_id, worldart_link, shikimori_id)
	Lgbt              *bool  `json:"lgbt,omitempty"`               // Фильтр по контенту LGBT (true/false)
	Sort              string `json:"sort,omitempty"`               // Сортировка по названию страны (title) или количеству материалов (count)

	// Фильтрация по внешним данным
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность (в минутах, точное значение или диапазон, например, 30 или 40-80)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (например, 7.0 или 6.5-8.2)
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг (например, 7.0 или 6.5-8.2)
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori (например, 7.0 или 6.5-8.2)
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList (например, 7.0 или 6.5-8.2)
	Actors            string `json:"actors,omitempty"`             // Список актёров через запятую
	Directors         string `json:"directors,omitempty"`          // Список режиссёров через запятую
	Producers         string `json:"producers,omitempty"`          // Список продюсеров через запятую
	Writers           string `json:"writers,omitempty"`            // Список сценаристов через запятую
	Composers         string `json:"composers,omitempty"`          // Список композиторов через запятую
	Editors           string `json:"editors,omitempty"`            // Список монтажёров через запятую
	Designers         string `json:"designers,omitempty"`          // Список дизайнеров через запятую
	Operators         string `json:"operators,omitempty"`          // Список операторов через запятую
	RatingMPAA        string `json:"rating_mpaa,omitempty"`        // Рейтинг MPAA (например, G, PG, PG-13, R и т.д.)
	MinimalAge        string `json:"minimal_age,omitempty"`        // Минимальный возраст (например, 16 или 12-16)
	AnimeKind         string `json:"anime_kind,omitempty"`         // Тип аниме (tv, movie, ova, ona, и т.д.)
	MydramalistTags   string `json:"mydramalist_tags,omitempty"`   // Теги MyDramaList через запятую (например, Friendship, Violence, etc.)
	AnimeStatus       string `json:"anime_status,omitempty"`       // Статус аниме (anons, ongoing, released, и т.д.)
	DramaStatus       string `json:"drama_status,omitempty"`       // Статус драмы (anons, ongoing, released, и т.д.)
	AllStatus         string `json:"all_status,omitempty"`         // Универсальный статус (анонс, ongoing, released и т.д.)
	AnimeStudios      string `json:"anime_studios,omitempty"`      // Студии для аниме через запятую (например, J.C.Staff, Studio Hibari)
	AnimeLicensedBy   string `json:"anime_licensed_by,omitempty"`  // Лицензионные правообладатели через запятую (например, Wakanim, Russian Reportage)
}

// ToMap преобразует структуру CountriesParams в карту параметров для HTTP-запроса.
func (cp *CountriesParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"types":              cp.Types,
		"year":               cp.Year,
		"translation_type":   cp.TranslationType,
		"has_field":          cp.HasField,
		"sort":               cp.Sort,
		"genres":             cp.Genres,
		"anime_genres":       cp.AnimeGenres,
		"drama_genres":       cp.DramaGenres,
		"all_genres":         cp.AllGenres,
		"duration":           cp.Duration,
		"kinopoisk_rating":   cp.KinopoiskRating,
		"imdb_rating":        cp.ImdbRating,
		"shikimori_rating":   cp.ShikimoriRating,
		"mydramalist_rating": cp.MydramalistRating,
		"actors":             cp.Actors,
		"directors":          cp.Directors,
		"producers":          cp.Producers,
		"writers":            cp.Writers,
		"composers":          cp.Composers,
		"editors":            cp.Editors,
		"designers":          cp.Designers,
		"operators":          cp.Operators,
		"rating_mpaa":        cp.RatingMPAA,
		"minimal_age":        cp.MinimalAge,
		"anime_kind":         cp.AnimeKind,
		"mydramalist_tags":   cp.MydramalistTags,
		"anime_status":       cp.AnimeStatus,
		"drama_status":       cp.DramaStatus,
		"all_status":         cp.AllStatus,
		"anime_studios":      cp.AnimeStudios,
		"anime_licensed_by":  cp.AnimeLicensedBy,
	}

	for k, v := range fields {
		if v != "" {
			params[k] = v
		}
	}

	if cp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(cp.TranslationID)
	}
	if cp.BlockTranslations != "" {
		params["block_translations"] = strings.ReplaceAll(cp.BlockTranslations, " ", "")
	}
	if cp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*cp.Lgbt)
	}

	return params
}
