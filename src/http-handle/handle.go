package handle

import (
	"api.example.com/pkg/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Services struct {
	User user.Server
}

func New(s *Services) http.Handler {
	mux := mux.NewRouter()

	func(user *userHandler) {
		mux.HandleFunc("/user", user.handleUsers)
		mux.HandleFunc("/user/{user_id}", user.handleUser)
	}(newUserHandler(s.User))

	return mux
}
