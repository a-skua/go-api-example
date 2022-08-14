package handle

import (
	"api.example.com/pkg/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Repository interface {
	user.Repository
}

func New(r Repository) http.Handler {
	mux := mux.NewRouter()

	{
		user := newUserHandler(user.NewServer(r))
		mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				user.create(w, r)
			default:
				http.NotFound(w, r)
			}
		})
		mux.HandleFunc("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
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
		})
	}

	return mux
}
