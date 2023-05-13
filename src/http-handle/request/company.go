package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"api.example.com/pkg/company"
	"github.com/gorilla/mux"
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

func parseCompanyPath(r *http.Request) (company.ID, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["company_id"])
	if err != nil {
		return 0, err
	}

	return company.ID(id), nil
}

func CompanyRead(req *http.Request) (company.ID, error) {
	id, err := parseCompanyPath(req)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.CompanyRead: %w", err)
	}

	return id, nil
}
