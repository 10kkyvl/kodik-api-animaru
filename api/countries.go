package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Countries выполняет запрос к эндпоинту /countries API Kodik с параметрами models.CountriesParams.
func Countries(c *client.Client, cp *models.CountriesParams) (*models.CountriesResponse, error) {
	var response models.CountriesResponse

	var params map[string]string
	if cp != nil {
		params = cp.ToMap()
	}
	err := c.DoRequest("GET", "/countries", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
