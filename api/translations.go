package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Translations выполняет запрос к эндпоинту /translations/v2 API Kodik,
// используя параметры, заданные в структуре models.TranslationsParams.
func Translations(c *client.Client, tp *models.TranslationsParams) (*models.TranslationsResponse, error) {
	var response models.TranslationsResponse
	var params map[string]string
	if tp != nil {
		params = tp.ToMap()
	}
	err := c.DoRequest("GET", "/translations/v2", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
