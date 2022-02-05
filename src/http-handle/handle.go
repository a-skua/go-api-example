package handle

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"api.example.com/pkg/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func New(r repository.User) http.Handler {
	mux := mux.NewRouter()

	user := &user{service.NewUser(r)}
	mux.HandleFunc("/user", user.create).Methods(http.MethodPost)
	mux.HandleFunc("/user/{user_id:[0-9]+}", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			user.read(w, req)
		case http.MethodPut:
			user.update(w, req)
		case http.MethodDelete:
			user.delete(w, req)
		}
	}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	return mux
}

// 認証情報の取得
func authUser(req *http.Request) (entity.UserID, error) {
	userID, err := strconv.Atoi(req.Header.Get("X-User-Id"))
	if err != nil {
		return 0, fmt.Errorf("handle.authUser: %w", err)
	}

	return entity.UserID(userID), nil
}
