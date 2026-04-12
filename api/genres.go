package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Genres выполняет запрос к эндпоинту /genres API Kodik с параметрами models.GenresParams.
func Genres(c *client.Client, gp *models.GenresParams) (*models.GenresResponse, error) {
	var response models.GenresResponse
	var params map[string]string
	if gp != nil {
		params = gp.ToMap()
	}
	err := c.DoRequest("GET", "/genres", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
