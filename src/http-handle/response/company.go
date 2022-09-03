package response

import (
	companies "api.example.com/pkg/company"
	"encoding/json"
	"net/http"
	"time"
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
