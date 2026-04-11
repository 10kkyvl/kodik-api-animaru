package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Years выполняет запрос к эндпоинту /years API Kodik с параметрами models.YearsParams.
func Years(yp *models.YearsParams) (*models.YearsResponse, error) {
	var response models.YearsResponse
	params := yp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/years", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
