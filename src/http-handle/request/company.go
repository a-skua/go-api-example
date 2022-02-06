package request

import (
	"api.example.com/pkg/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func CompanyCreate(req *http.Request) (*entity.Company, error) {
	body := struct {
		Company struct {
			Name string `json:"name"`
		} `json:"company"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("HTTP Request CompanyCreate: %w", err)
	}

	return entity.NewCompany(
		body.Company.Name,
	), nil
}
