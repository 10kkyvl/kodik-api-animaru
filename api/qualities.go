package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Qualities выполняет запрос к эндпоинту /qualities/v2 API Kodik, используя параметры models.QualitiesParams.
func Qualities(c *client.Client, qp *models.QualitiesParams) (*models.QualitiesResponse, error) {
	var response models.QualitiesResponse
	var params map[string]string
	if qp != nil {
		params = qp.ToMap()
	}
	err := c.DoRequest("GET", "/qualities/v2", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
