package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Search выполняет запрос к эндпоинту /search API Kodik, используя параметры, заданные в структуре models.SearchParams.
func Search(c *client.Client, sp *models.SearchParams) (*models.SearchResponse, error) {
	var response models.SearchResponse
	var params map[string]string
	if sp != nil {
		params = sp.ToMap()
	}
	err := c.DoRequest("GET", "/search", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
