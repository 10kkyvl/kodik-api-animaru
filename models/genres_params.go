package models

import (
	"strconv"
	"strings"
)

// GenresParams описывает параметры запроса для эндпоинта /genres API Kodik.
// Токен не передаётся, так как он уже установлен в клиенте.
type GenresParams struct {
	// Основные параметры
	GenresType        string `json:"genres_type,omitempty"`        // Выбор источника жанров: kinopoisk (по умолчанию), shikimori, mydramalist или all
	Types             string `json:"types,omitempty"`              // Фильтрация по типу материала (например, foreign-movie, cartoon-serial)
	Year              string `json:"year,omitempty"`               // Фильтр по году или диапазону лет
	TranslationID     int    `json:"translation_id,omitempty"`     // Фильтр по ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // Исключение указанных озвучек (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода: voice или subtitles
	HasField          string `json:"has_field,omitempty"`          // Фильтрация по наличию определённых полей (например, kinopoisk_id, imdb_id, mdl_id, worldart_link, shikimori_id)
	Lgbt              *bool  `json:"lgbt,omitempty"`               // Фильтрация по содержанию LGBT (true/false)
	Sort              string `json:"sort,omitempty"`               // Сортировка результатов по названию жанра (title) или по числу материалов (count)

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Фильтр по странам (через запятую)
	Genres            string `json:"genres,omitempty"`             // Фильтр по жанрам (через запятую)
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Фильтр по жанрам для аниме
	DramaGenres       string `json:"drama_genres,omitempty"`       // Фильтр по жанрам для драм
	AllGenres         string `json:"all_genres,omitempty"`         // Фильтр по всем жанрам
	Duration          string `json:"duration,omitempty"`           // Длительность (точное значение или диапазон в минутах)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (точное значение или диапазон)
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг (точное значение или диапазон)
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori (точное значение или диапазон)
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList (точное значение или диапазон)
	Actors            string `json:"actors,omitempty"`             // Список актёров через запятую
	Directors         string `json:"directors,omitempty"`          // Список режиссёров через запятую
	Producers         string `json:"producers,omitempty"`          // Список продюсеров через запятую
	Writers           string `json:"writers,omitempty"`            // Список сценаристов через запятую
	Composers         string `json:"composers,omitempty"`          // Список композиторов через запятую
	Editors           string `json:"editors,omitempty"`            // Список монтажёров через запятую
	Designers         string `json:"designers,omitempty"`          // Список дизайнеров через запятую
	Operators         string `json:"operators,omitempty"`          // Список операторов через запятую
	RatingMPAA        string `json:"rating_mpaa,omitempty"`        // Рейтинг MPAA (например, G, PG, PG-13, R)
	MinimalAge        string `json:"minimal_age,omitempty"`        // Минимальный возраст для просмотра (например, 16 или 12-16)
	AnimeKind         string `json:"anime_kind,omitempty"`         // Тип аниме: tv, movie, ova, ona и т.д.
	MydramalistTags   string `json:"mydramalist_tags,omitempty"`   // Теги MyDramaList через запятую
	AnimeStatus       string `json:"anime_status,omitempty"`       // Статус аниме (anons, ongoing, released и т.д.)
	DramaStatus       string `json:"drama_status,omitempty"`       // Статус драмы (anons, ongoing, released и т.д.)
	AllStatus         string `json:"all_status,omitempty"`         // Универсальный статус (например, anons, ongoing, released)
	AnimeStudios      string `json:"anime_studios,omitempty"`      // Студии для аниме через запятую
	AnimeLicensedBy   string `json:"anime_licensed_by,omitempty"`  // Правообладатели для аниме через запятую
}

// ToMap преобразует структуру GenresParams в карту параметров для формирования HTTP-запроса.
func (gp *GenresParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"genres_type":        gp.GenresType,
		"types":              gp.Types,
		"year":               gp.Year,
		"translation_type":   gp.TranslationType,
		"has_field":          gp.HasField,
		"sort":               gp.Sort,
		"countries":          gp.Countries,
		"genres":             gp.Genres,
		"anime_genres":       gp.AnimeGenres,
		"drama_genres":       gp.DramaGenres,
		"all_genres":         gp.AllGenres,
		"duration":           gp.Duration,
		"kinopoisk_rating":   gp.KinopoiskRating,
		"imdb_rating":        gp.ImdbRating,
		"shikimori_rating":   gp.ShikimoriRating,
		"mydramalist_rating": gp.MydramalistRating,
		"actors":             gp.Actors,
		"directors":          gp.Directors,
		"producers":          gp.Producers,
		"writers":            gp.Writers,
		"composers":          gp.Composers,
		"editors":            gp.Editors,
		"designers":          gp.Designers,
		"operators":          gp.Operators,
		"rating_mpaa":        gp.RatingMPAA,
		"minimal_age":        gp.MinimalAge,
		"anime_kind":         gp.AnimeKind,
		"mydramalist_tags":   gp.MydramalistTags,
		"anime_status":       gp.AnimeStatus,
		"drama_status":       gp.DramaStatus,
		"all_status":         gp.AllStatus,
		"anime_studios":      gp.AnimeStudios,
		"anime_licensed_by":  gp.AnimeLicensedBy,
	}

	for k, v := range fields {
		if v != "" {
			params[k] = v
		}
	}

	if gp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(gp.TranslationID)
	}
	if gp.BlockTranslations != "" {
		params["block_translations"] = strings.ReplaceAll(gp.BlockTranslations, " ", "")
	}
	if gp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*gp.Lgbt)
	}

	return params
}
