package request

import (
	"api.example.com/pkg/user"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func userValue(req *http.Request) (*user.User, error) {
	defer req.Body.Close()
	body := struct {
		User struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.userValue: %w", err)
	}

	password, err := user.NewPassword([]byte(body.User.Password))
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.userValue: %w", err)
	}

	return user.New(
		body.User.Name,
		password,
	), nil
}

func userKey(req *http.Request) (user.ID, error) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.userKey: %w", err)
	}

	return user.ID(id), nil
}

func UserCreate(req *http.Request) (*user.User, error) {
	user, err := userValue(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserCreate: %w", err)
	}
	return user, nil
}

func UserRead(req *http.Request) (user.ID, error) {
	userID, err := userKey(req)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.UserRead: %w", err)
	}
	return userID, nil
}

func UserUpdate(req *http.Request) (*user.User, error) {
	userID, err := userKey(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserUpdate: %w", err)
	}

	user, err := userValue(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserUpdate: %w", err)
	}

	user.ID = userID
	return user, nil
}

func UserDelete(req *http.Request) (user.ID, error) {
	userID, err := userKey(req)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.UserDelete: %w", err)
	}
	return userID, nil
}
