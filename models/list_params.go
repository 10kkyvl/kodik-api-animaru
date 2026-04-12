package models

import (
	"strconv"
	"strings"
)

// ListParams описывает доступные параметры для запроса /list API Kodik.
// Если поле имеет нулевое значение, оно не включается в запрос.
type ListParams struct {
	// Основные параметры
	Limit             int    `json:"limit,omitempty"`              // Количество материалов на запрос (1-100)
	Sort              string `json:"sort,omitempty"`               // Поле сортировки (например, updated_at)
	Order             string `json:"order,omitempty"`              // Направление сортировки (asc или desc)
	Types             string `json:"types,omitempty"`              // Список типов материалов через запятую
	Year              string `json:"year,omitempty"`               // Год или диапазон годов выпуска
	TranslationID     int    `json:"translation_id,omitempty"`     // ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // IDs озвучек, которые следует исключить (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода: voice или subtitles
	HasField          string `json:"has_field,omitempty"`          // Фильтрация по наличию поля

	// Флаги маркеров
	Camrip *bool `json:"camrip,omitempty"` // Показывать только камрипы (true/false)
	Lgbt   *bool `json:"lgbt,omitempty"`   // Фильтрация материалов по содержанию LGBT-сцен

	// Параметры для сериалов
	WithSeasons      bool   `json:"with_seasons,omitempty"`       // Включать данные о сезонах
	WithEpisodes     bool   `json:"with_episodes,omitempty"`      // Включать данные об эпизодах
	WithEpisodesData bool   `json:"with_episodes_data,omitempty"` // Включать подробные данные об эпизодах
	WithPageLinks    bool   `json:"with_page_links,omitempty"`    // Заменять ссылки на плеер на ссылки на страницу
	NotBlockedIn     string `json:"not_blocked_in,omitempty"`     // Список стран (через запятую), в которых материал не должен быть заблокирован
	NotBlockedForMe  *bool  `json:"not_blocked_for_me,omitempty"` // Флаг исключения заблокированных материалов (автоматическое определение страны)

	// Внешняя фильтрация данных
	WithMaterialData bool   `json:"with_material_data,omitempty"` // Включать расширенные данные материала
	Countries        string `json:"countries,omitempty"`          // Страны через запятую
	Genres           string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres      string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres      string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres        string `json:"all_genres,omitempty"`         // Все жанры через запятую

	// Дополнительные параметры
	Duration          string `json:"duration,omitempty"`           // Длительность (точное значение или диапазон)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг по Кинопоиску или диапазон
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг или диапазон
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Шикимори или диапазон
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
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме (tv, movie, ova, ona и т.д.)
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме (anons, ongoing, released и т.д.)
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Владелец лицензионных прав через запятую
}

// ToMap преобразует структуру ListParams в карту параметров для HTTP-запроса.
func (lp *ListParams) ToMap() map[string]string {
	params := make(map[string]string)

	fields := map[string]string{
		"sort":               lp.Sort,
		"order":              lp.Order,
		"types":              lp.Types,
		"year":               lp.Year,
		"block_translations": lp.BlockTranslations,
		"translation_type":   lp.TranslationType,
		"has_field":          lp.HasField,
		"not_blocked_in":     lp.NotBlockedIn,
		"countries":          lp.Countries,
		"genres":             lp.Genres,
		"anime_genres":       lp.AnimeGenres,
		"drama_genres":       lp.DramaGenres,
		"all_genres":         lp.AllGenres,
		"duration":           lp.Duration,
		"kinopoisk_rating":   lp.KinopoiskRating,
		"imdb_rating":        lp.ImdbRating,
		"shikimori_rating":   lp.ShikimoriRating,
		"mydramalist_rating": lp.MydramalistRating,
		"actors":             lp.Actors,
		"directors":          lp.Directors,
		"producers":          lp.Producers,
		"writers":            lp.Writers,
		"composers":          lp.Composers,
		"editors":            lp.Editors,
		"designers":          lp.Designers,
		"operators":          lp.Operators,
		"rating_mpaa":        lp.RatingMPAA,
		"minimal_age":        lp.MinimalAge,
		"anime_kind":         lp.AnimeKind,
		"mydramalist_tags":   lp.MydramalistTags,
		"anime_status":       lp.AnimeStatus,
		"drama_status":       lp.DramaStatus,
		"all_status":         lp.AllStatus,
		"anime_studios":      lp.AnimeStudios,
		"anime_licensed_by":  lp.AnimeLicensedBy,
	}

	for k, v := range fields {
		if v != "" {
			if k == "not_blocked_in" || k == "block_translations" {
				params[k] = strings.ReplaceAll(v, " ", "")
			} else {
				params[k] = v
			}
		}
	}

	if lp.Limit != 0 {
		params["limit"] = strconv.Itoa(lp.Limit)
	}
	if lp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(lp.TranslationID)
	}

	// Булевые значения
	if lp.Camrip != nil {
		params["camrip"] = strconv.FormatBool(*lp.Camrip)
	}
	if lp.Lgbt != nil {
		params["lgbt"] = strconv.FormatBool(*lp.Lgbt)
	}
	if lp.NotBlockedForMe != nil {
		params["not_blocked_for_me"] = strconv.FormatBool(*lp.NotBlockedForMe)
	}

	// Флаги (явная передача)
	params["with_seasons"] = strconv.FormatBool(lp.WithSeasons)
	params["with_episodes"] = strconv.FormatBool(lp.WithEpisodes)
	params["with_episodes_data"] = strconv.FormatBool(lp.WithEpisodesData)
	params["with_page_links"] = strconv.FormatBool(lp.WithPageLinks)
	params["with_material_data"] = strconv.FormatBool(lp.WithMaterialData)

	return params
}
