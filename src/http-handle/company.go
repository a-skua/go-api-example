package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/company"
	"log"
	"net/http"
)

type companyHandler struct {
	server company.Server
}

func (h *companyHandler) create(w http.ResponseWriter, r *http.Request) {
	company, err := request.CompanyCreate(r)
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

	err = response.CompanyCreate(w, company)
	if err != nil {
		log.Println(err)
	}
}

func (h *companyHandler) read(w http.ResponseWriter, r *http.Request) {
	id, err := request.CompanyRead(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	company, err := h.server.Read(id)
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

func (h *companyHandler) update(w http.ResponseWriter, r *http.Request) {
	company, err := request.CompanyUpdate(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	company, err = h.server.Update(company)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.CompanyUpdate(w, company)
	if err != nil {
		log.Println(err)
	}
}

func (h *companyHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.CompanyDelete(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = h.server.Delete(id)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.CompanyDelete(w)
	if err != nil {
		log.Println(err)
	}
}
