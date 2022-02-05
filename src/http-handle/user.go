package handle

import (
	"api.example.com/http-handle/request"
	"api.example.com/http-handle/response"
	"api.example.com/pkg/service"
	"fmt"
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
		err := response.Error(w, http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	user, err = u.service.Create(user)
	if err != nil {
		log.Println(err)
		err := response.Error(w, http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = response.UserCreate(w, user)
	if err != nil {
		log.Println(err)
	}
}

func (user) read(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(w, "read user")
}

func (user) update(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(w, "update user")
}

func (user) delete(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(w, "delete user")
}
