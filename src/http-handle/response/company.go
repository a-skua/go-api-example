package response

import (
	"api.example.com/pkg/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func CompanyCreate(w http.ResponseWriter, company *entity.Company) error {
	type Company struct {
		ID   entity.CompanyID `json:"id"`
		Name string           `json:"name"`
	}

	res := struct {
		Company Company `json:"company"`
	}{
		Company: Company{
			ID:   company.ID,
			Name: company.Name,
		},
	}

	writeHeader(w)
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		return fmt.Errorf("HTTP Response CompanyCreate: %w", err)
	}
	return nil
}
