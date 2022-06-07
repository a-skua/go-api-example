package request

import (
	"api.example.com/pkg/company"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func parseCompanyBody(r *http.Request) (*company.Company, error) {
	defer r.Body.Close()

	body := struct {
		Company struct {
			Name company.Name `json:"name"`
		} `json:"company"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return company.New(
		body.Company.Name,
	), nil
}

func parseCompanyPath(r *http.Request) (company.ID, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["company_id"])
	return company.ID(id), err
}

func CompanyCreate(r *http.Request) (*company.Company, error) {
	company, err := parseCompanyBody(r)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.CompanyCreate: %w", err)
	}

	return company, nil
}

func CompanyRead(r *http.Request) (company.ID, error) {
	id, err := parseCompanyPath(r)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.CompanyRead: %w", err)
	}

	return id, nil
}

func CompanyUpdate(r *http.Request) (*company.Company, error) {
	id, err := parseCompanyPath(r)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.CompanyUpdate: %w", err)
	}

	company, err := parseCompanyBody(r)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.CompanyUpdate: %w", err)
	}

	company.ID = id
	return company, nil
}

func CompanyDelete(r *http.Request) (company.ID, error) {
	id, err := parseCompanyPath(r)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.CompanyDelete: %w", err)
	}

	return id, nil
}
