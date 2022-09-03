package handle

import (
	"api.example.com/pkg/company"
	"api.example.com/pkg/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Services struct {
	User    user.Server
	Company company.Server
}

func New(s *Services) http.Handler {
	mux := mux.NewRouter()

	func(user *userHandler) {
		mux.HandleFunc("/user", user.handleUsers)
		mux.HandleFunc("/user/{user_id}", user.handleUser)
	}(newUserHandler(s.User))

	func(company *companyHandler) {
		mux.HandleFunc("/company", company.handleCompanies)
	}(newCompanyHandler(s.Company))

	return mux
}
