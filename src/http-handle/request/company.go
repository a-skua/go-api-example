package request

import (
	"api.example.com/pkg/company"
	"encoding/json"
	"net/http"
)

func NewCompanyCreate(r *http.Request) (*company.Company, error) {
	defer r.Body.Close()

	body := struct {
		Company struct {
			Name    company.Name    `json:"name"`
			OwnerID company.OwnerID `json:"owner_id"`
		}
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return company.New(body.Company.Name, body.Company.OwnerID), nil
}
