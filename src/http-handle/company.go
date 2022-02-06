package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/service"
	"log"
	"net/http"
)

type company struct {
	service *service.Company
}

func (c *company) create(w http.ResponseWriter, req *http.Request) {
	company, err := request.CompanyCreate(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	authID, err := authUser(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	company, err = c.service.Create(company, authID)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	err = response.CompanyCreate(w, company)
	if err != nil {
		log.Println(err)
	}
}
