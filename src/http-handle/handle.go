package handle

import (
	"api.example.com/http-handle/response"
	"api.example.com/pkg/user"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Repository interface {
	user.Repository
}

func New(r Repository) http.Handler {
	mux := mux.NewRouter()

	{
		user := &userHandle{user.NewService(r)}
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
	}
	return mux
}

// 認証情報の取得
type auth user.ID

func newAuth(req *http.Request) (auth, error) {
	userID, err := strconv.Atoi(req.Header.Get("X-User-Id"))
	if err != nil {
		return 0, fmt.Errorf("handle.authUser: %w", err)
	}

	return auth(userID), nil
}

func (a auth) verify(userid user.ID) bool {
	return a == auth(userid)
}

func writeError(w http.ResponseWriter, statusCode int) {
	err := response.Error(w, http.StatusBadRequest)
	if err != nil {
		log.Println(err)
	}
}
