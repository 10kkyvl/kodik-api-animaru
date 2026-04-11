package api

import (
	"github.com/10kkyvl/kodik-api-animaru/client"
	"github.com/10kkyvl/kodik-api-animaru/models"
)

// List выполняет запрос к эндпоинту /list API Kodik, используя параметры, заданные в структуре models.ListParams.
func List(lp *models.ListParams) (*models.ListResponse, error) {
	var response models.ListResponse
	params := lp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/list", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
