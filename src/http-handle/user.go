package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/user"
	"log"
	"net/http"
)

type userHandle struct {
	service user.Server
}

func (h *userHandle) create(w http.ResponseWriter, req *http.Request) {
	user, err := request.UserCreate(req)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err = h.service.Create(user)
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

func (h *userHandle) read(w http.ResponseWriter, req *http.Request) {
	userID, err := request.UserRead(req)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err := h.service.Read(userID)
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

func (h *userHandle) update(w http.ResponseWriter, req *http.Request) {
	user, err := request.UserUpdate(req)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	user, err = h.service.Update(user)
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

func (h *userHandle) delete(w http.ResponseWriter, req *http.Request) {
	userID, err := request.UserDelete(req)
	if err != nil {
		log.Println(err)
		response.Error(w, err)
		return
	}

	err = h.service.Delete(userID)
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
