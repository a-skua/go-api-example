package handle

import (
	"log"
	"net/http"

	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/company"
)

type companyHandler struct {
	server company.Server
}

func newCompanyHandler(s company.Server) *companyHandler {
	return &companyHandler{s}
}

func (h *companyHandler) handleCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *companyHandler) handleCompany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *companyHandler) create(w http.ResponseWriter, r *http.Request) {
	company, err := request.NewCompanyCreate(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	company, err = h.server.Create(company)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.WriteCompany(w, company)
	if err != nil {
	}
}

func (h *companyHandler) read(w http.ResponseWriter, r *http.Request) {
	companyId, err := request.CompanyRead(r)

	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	company, err := h.server.Read(companyId)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.CompanyRead(w, company)
	if err != nil {
		log.Println(err)
	}
}
