package response

import (
	"api.example.com/pkg/company"
	"encoding/json"
	"fmt"
	"net/http"
)

func writeCompany(w http.ResponseWriter, c *company.Company) error {
	type Company struct {
		ID   company.ID   `json:"id"`
		Name company.Name `json:"name"`
	}

	body := struct {
		Company `json:"company"`
	}{
		Company{c.ID, c.Name},
	}

	writeHeader(w)
	return json.NewEncoder(w).Encode(&body)
}

func CompanyCreate(w http.ResponseWriter, c *company.Company) error {
	err := writeCompany(w, c)
	if err != nil {
		return fmt.Errorf("http-handle/response.CompanyCreate: %w", err)
	}

	return nil
}

func CompanyRead(w http.ResponseWriter, c *company.Company) error {
	err := writeCompany(w, c)
	if err != nil {
		return fmt.Errorf("http-handle/response.CompanyRead: %w", err)
	}

	return nil
}

func CompanyUpdate(w http.ResponseWriter, c *company.Company) error {
	err := writeCompany(w, c)
	if err != nil {
		return fmt.Errorf("http-handle/response.CompanyUpdate: %w", err)
	}

	return nil
}

func CompanyDelete(w http.ResponseWriter) error {
	body := struct {
		Company struct{} `json:"company"`
	}{}

	writeHeader(w)
	err := json.NewEncoder(w).Encode(&body)
	if err != nil {
		return fmt.Errorf("http-handle/response.CompanyDelete: %w", err)
	}

	return nil
}
