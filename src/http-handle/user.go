package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/service"
	"log"
	"net/http"
)

type user struct {
	service *service.User
}

func (u *user) create(w http.ResponseWriter, req *http.Request) {
	user, err := request.UserCreate(req)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	user, err = u.service.Create(user)
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

func (u *user) read(w http.ResponseWriter, req *http.Request) {
	userID, err := request.UserRead(req)
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

	user, err := u.service.Read(userID, authID)
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

func (u *user) update(w http.ResponseWriter, req *http.Request) {
	user, err := request.UserUpdate(req)
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

	user, err = u.service.Update(user, authID)
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

func (u *user) delete(w http.ResponseWriter, req *http.Request) {
	userID, err := request.UserDelete(req)
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

	err = u.service.Delete(userID, authID)
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
