package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api.example.com/pkg/company"
	companies "api.example.com/pkg/company"
)

func WriteCompany(w http.ResponseWriter, company *companies.Company) error {
	type value struct {
		ID        companies.ID      `json:"id"`
		Name      companies.Name    `json:"name"`
		OwnerID   companies.OwnerID `json:"owner_id"`
		UpdatedAt time.Time         `json:"updated_at"`
	}

	body := struct {
		Company value `json:"company"`
	}{
		Company: value{
			ID:        company.ID,
			Name:      company.Name,
			OwnerID:   company.OwnerID,
			UpdatedAt: company.UpdatedAt,
		},
	}

	writeHeader(w)
	return json.NewEncoder(w).Encode(&body)
}

func CompanyRead(w http.ResponseWriter, c *company.Company) error {
	err := WriteCompany(w, c)
	if err != nil {
		return fmt.Errorf("http-handle/response.CompanyRead: %w", err)
	}

	return nil
}
