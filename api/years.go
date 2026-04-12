package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// Years выполняет запрос к эндпоинту /years API Kodik с параметрами models.YearsParams.
func Years(c *client.Client, yp *models.YearsParams) (*models.YearsResponse, error) {
	var response models.YearsResponse
	var params map[string]string
	if yp != nil {
		params = yp.ToMap()
	}
	err := c.DoRequest("GET", "/years", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
