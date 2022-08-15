package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/user"
	"log"
	"net/http"
)

type userHandler struct {
	server user.Server
}

func newUserHandler(s user.Server) *userHandler {
	return &userHandler{s}
}

func (h *userHandler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *userHandler) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.read(w, r)
	case http.MethodPut:
		h.update(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *userHandler) create(w http.ResponseWriter, r *http.Request) {
	user, err := request.UserCreate(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err = h.server.Create(user)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.UserCreate(w, user)
	if err != nil {
		log.Println(err)
	}
}

func (h *userHandler) read(w http.ResponseWriter, r *http.Request) {
	userID, err := request.UserRead(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err := h.server.Read(userID)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.UserRead(w, user)
	if err != nil {
		log.Println(err)
	}
}

func (h *userHandler) update(w http.ResponseWriter, r *http.Request) {
	user, err := request.UserUpdate(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err = h.server.Update(user)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.UserUpdate(w, user)
	if err != nil {
		log.Println(err)
	}
}

func (h *userHandler) delete(w http.ResponseWriter, r *http.Request) {
	userID, err := request.UserDelete(r)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = h.server.Delete(userID)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = response.UserDelete(w)
	if err != nil {
		log.Println(err)
	}
}
