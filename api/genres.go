package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Genres выполняет запрос к эндпоинту /genres API Kodik с параметрами models.GenresParams.
func Genres(gp *models.GenresParams) (*models.GenresResponse, error) {
	var response models.GenresResponse
	params := gp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/genres", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
