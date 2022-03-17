package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/user"
	"log"
	"net/http"
)

type userHandle struct {
	service user.Service
}

func (h *userHandle) create(w http.ResponseWriter, req *http.Request) {
	user, err := request.UserCreate(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	user, err = h.service.Create(user)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
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
		writeError(w, http.StatusBadRequest)
		return
	}

	auth, err := newAuth(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}
	ok := auth.verify(userID)
	if !ok {
		log.Printf("unauthorized user-id=%v.\n", userID)
		writeError(w, http.StatusForbidden)
		return
	}

	user, err := h.service.Read(userID)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
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
		writeError(w, http.StatusBadRequest)
		return
	}

	auth, err := newAuth(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	ok := auth.verify(user.ID)
	if !ok {
		log.Printf("unauthorized user-id=%v.\n", user.ID)
		writeError(w, http.StatusForbidden)
		return
	}

	user, err = h.service.Update(user)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
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
		writeError(w, http.StatusBadRequest)
		return
	}

	auth, err := newAuth(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	ok := auth.verify(userID)
	if !ok {
		log.Printf("unauthorized user-id=%v.\n", userID)
		writeError(w, http.StatusForbidden)
		return
	}

	err = h.service.Delete(userID)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	err = response.UserDelete(w)
	if err != nil {
		log.Println(err)
	}
}
