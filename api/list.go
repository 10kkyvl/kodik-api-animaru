package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// List выполняет запрос к эндпоинту /list API Kodik, используя параметры, заданные в структуре models.ListParams.
func List(c *client.Client, lp *models.ListParams) (*models.ListResponse, error) {
	var response models.ListResponse
	var params map[string]string
	if lp != nil {
		params = lp.ToMap()
	}
	err := c.DoRequest("GET", "/list", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
