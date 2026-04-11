package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Qualities выполняет запрос к эндпоинту /qualities/v2 API Kodik, используя параметры models.QualitiesParams.
func Qualities(qp *models.QualitiesParams) (*models.QualitiesResponse, error) {
	var response models.QualitiesResponse
	params := qp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/qualities/v2", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
