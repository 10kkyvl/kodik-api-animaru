package models

import (
	"strconv"
	"strings"
)

// SearchParams описывает параметры запроса для эндпоинта /search API Kodik.
// Поля соответствуют описанным в таблице параметров.
// Если значение поля не установлено (нулевое), оно не будет включено в итоговый запрос.
type SearchParams struct {
	// Обязательные поля (хотя бы одно из них должно быть заполнено)
	Title     string `json:"title,omitempty"`
	TitleOrig string `json:"title_orig,omitempty"`

	// Параметры строгого поиска
	Strict    bool `json:"strict,omitempty"`
	FullMatch bool `json:"full_match,omitempty"`

	// Поиск по идентификаторам
	ID                  string `json:"id,omitempty"`
	PlayerLink          string `json:"player_link,omitempty"`
	KinopoiskID         int    `json:"kinopoisk_id,omitempty"`
	ImdbID              string `json:"imdb_id,omitempty"`
	MdlID               string `json:"mdl_id,omitempty"`
	WorldartAnimationID int    `json:"worldart_animation_id,omitempty"`
	WorldartCinemaID    int    `json:"worldart_cinema_id,omitempty"`
	WorldartLink        string `json:"worldart_link,omitempty"`
	ShikimoriID         int    `json:"shikimori_id,omitempty"`

	// Опциональные параметры
	Limit                     int    `json:"limit,omitempty"`
	Types                     string `json:"types,omitempty"` // список типов через запятую
	Year                      string `json:"year,omitempty"`
	TranslationID             int    `json:"translation_id,omitempty"`
	TranslationType           string `json:"translation_type,omitempty"`
	HasField                  string `json:"has_field,omitempty"`
	PrioritizeTranslations    string `json:"prioritize_translations,omitempty"`
	UnprioritizeTranslations  string `json:"unprioritize_translations,omitempty"`
	PrioritizeTranslationType string `json:"prioritize_translation_type,omitempty"`
	BlockTranslations         string `json:"block_translations,omitempty"`
	Camrip                    *bool  `json:"camrip,omitempty"` // указатель для отличия false и не заданного значения
	Lgbt                      *bool  `json:"lgbt,omitempty"`   // указатель для отличия false и не заданного значения
	WithSeasons               bool   `json:"with_seasons,omitempty"`
	Season                    int    `json:"season,omitempty"`
	WithEpisodes              bool   `json:"with_episodes,omitempty"`
	WithEpisodesData          bool   `json:"with_episodes_data,omitempty`
	Episode                   int    `json:"episode,omitempty"`
	WithPageLinks             bool   `json:"with_page_links,omitempty"`
	NotBlockedIn              string `json:"not_blocked_in,omitempty"` // список стран через запятую (без пробелов)
	NotBlockedForMe           *bool  `json:"not_blocked_for_me,omitempty"`
	WithMaterialData          bool   `json:"with_material_data,omitempty"`
	Countries                 string `json:"countries,omitempty"` // список стран через запятую
	Genres                    string `json:"genres,omitempty"`    // список жанров через запятую
	Duration                  string `json:"duration,omitempty"`
	KinopoiskRating           string `json:"kinopoisk_rating,omitempty"`
	ImdbRating                string `json:"imdb_rating,omitempty"`
	ShikimoriRating           string `json:"shikimori_rating,omitempty"`
	MydramalistRating         string `json:"mydramalist_rating,omitempty"`
	Actors                    string `json:"actors,omitempty"`    // список актёров через запятую
	Directors                 string `json:"directors,omitempty"` // список режиссёров через запятую
	Producers                 string `json:"producers,omitempty"` // список продюсеров через запятую
	Writers                   string `json:"writers,omitempty"`   // список сценаристов через запятую
	Composers                 string `json:"composers,omitempty"` // список композиторов через запятую
	Editors                   string `json:"editors,omitempty"`   // список монтажёров через запятую
	Designers                 string `json:"designers,omitempty"` // список дизайнеров через запятую
	Operators                 string `json:"operators,omitempty"` // список операторов через запятую
	RatingMPAA                string `json:"rating_mpaa,omitempty"`
	MinimalAge                string `json:"minimal_age,omitempty"`
	AnimeKind                 string `json:"anime_kind,omitempty"`
	MydramalistTags           string `json:"mydramalist_tags,omitempty"`
	AnimeStatus               string `json:"anime_status,omitempty"`
	DramaStatus               string `json:"drama_status,omitempty"`
	AllStatus                 string `json:"all_status,omitempty"`
	AnimeStudios              string `json:"anime_studios,omitempty"`
	AnimeLicensedBy           string `json:"anime_licensed_by,omitempty"`
}

// ToMap преобразует структуру SearchParams в карту параметров для запроса.
// Поля со значением по умолчанию пропускаются (за исключением булевых, для которых значение false явно передаётся).
func (sp *SearchParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"title":                        sp.Title,
		"title_orig":                   sp.TitleOrig,
		"id":                           sp.ID,
		"player_link":                  sp.PlayerLink,
		"imdb_id":                      sp.ImdbID,
		"mdl_id":                       sp.MdlID,
		"worldart_link":                sp.WorldartLink,
		"types":                        sp.Types,
		"year":                         sp.Year,
		"translation_type":             sp.TranslationType,
		"has_field":                    sp.HasField,
		"prioritize_translations":      sp.PrioritizeTranslations,
		"unprioritize_translations":    sp.UnprioritizeTranslations,
		"prioritize_translation_type":  sp.PrioritizeTranslationType,
		"block_translations":           sp.BlockTranslations,
		"not_blocked_in":               sp.NotBlockedIn,
		"countries":                    sp.Countries,
		"genres":                       sp.Genres,
		"duration":                     sp.Duration,
		"kinopoisk_rating":             sp.KinopoiskRating,
		"imdb_rating":                  sp.ImdbRating,
		"shikimori_rating":             sp.ShikimoriRating,
		"mydramalist_rating":           sp.MydramalistRating,
		"actors":                       sp.Actors,
		"directors":                    sp.Directors,
		"producers":                    sp.Producers,
		"writers":                      sp.Writers,
		"composers":                    sp.Composers,
		"editors":                      sp.Editors,
		"designers":                    sp.Designers,
		"operators":                    sp.Operators,
		"rating_mpaa":                  sp.RatingMPAA,
		"minimal_age":                  sp.MinimalAge,
		"anime_kind":                   sp.AnimeKind,
		"mydramalist_tags":             sp.MydramalistTags,
		"anime_status":                 sp.AnimeStatus,
		"drama_status":                 sp.DramaStatus,
		"all_status":                   sp.AllStatus,
		"anime_studios":                sp.AnimeStudios,
		"anime_licensed_by":            sp.AnimeLicensedBy,
	}

	for k, v := range fields {
		if v != "" {
			if k == "not_blocked_in" {
				params[k] = strings.ReplaceAll(v, " ", "")
			} else {
				params[k] = v
			}
		}
	}

	// Целочисленные поля
	intFields := map[string]int{
		"kinopoisk_id":           sp.KinopoiskID,
		"worldart_animation_id": sp.WorldartAnimationID,
		"worldart_cinema_id":    sp.WorldartCinemaID,
		"shikimori_id":          sp.ShikimoriID,
		"limit":                 sp.Limit,
		"translation_id":        sp.TranslationID,
		"season":                sp.Season,
		"episode":               sp.Episode,
	}

	for k, v := range intFields {
		if v != 0 {
			params[k] = strconv.Itoa(v)
		}
	}

	// Булевые значения (явная передача)
	params["strict"] = strconv.FormatBool(sp.Strict)
	params["full_match"] = strconv.FormatBool(sp.FullMatch)
	params["with_seasons"] = strconv.FormatBool(sp.WithSeasons)
	params["with_episodes"] = strconv.FormatBool(sp.WithEpisodes)
	params["with_episodes_data"] = strconv.FormatBool(sp.WithEpisodesData)
	params["with_page_links"] = strconv.FormatBool(sp.WithPageLinks)
	params["with_material_data"] = strconv.FormatBool(sp.WithMaterialData)

	// Указатели на булевые значения
	if sp.Camrip != nil {
		params["camrip"] = strconv.FormatBool(*sp.Camrip)
	}
	if sp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*sp.Lgbt)
	}
	if sp.NotBlockedForMe != nil {
		params["not_blocked_for_me"] = strconv.FormatBool(*sp.NotBlockedForMe)
	}

	return params
}
