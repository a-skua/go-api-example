package handle

import (
	"api.example.com/pkg/company"
	"api.example.com/pkg/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Repository interface {
	user.Repository
	company.Repository
}

func New(r Repository) http.Handler {
	mux := mux.NewRouter()

	{
		user := &userHandler{user.NewServer(r)}
		mux.HandleFunc("/user", user.create).Methods(http.MethodPost)
		mux.HandleFunc("/user/{user_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				user.read(w, r)
			case http.MethodPut:
				user.update(w, r)
			case http.MethodDelete:
				user.delete(w, r)
			default:
				http.NotFound(w, r)
			}
		}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	}
	{
		company := &companyHandler{company.NewServer(r)}
		mux.HandleFunc("/company", company.create).Methods(http.MethodPost)
		mux.HandleFunc("/company/{company_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				company.read(w, r)
			case http.MethodPut:
				company.update(w, r)
			case http.MethodDelete:
				company.delete(w, r)
			default:
				http.NotFound(w, r)
			}
		}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	}
	return mux
}
