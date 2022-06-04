package request

import (
	"api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func parseUserValue(r *http.Request) (*user.User, error) {
	defer r.Body.Close()

	body := struct {
		User struct {
			Name     user.Name `json:"name"`
			Password string    `json:"password"`
		} `json:"user"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	password, err := password.New(body.User.Password)
	if err != nil {
		return nil, err
	}

	return user.New(
		body.User.Name,
		password,
	), nil
}

func parseUserID(r *http.Request) (user.ID, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return 0, err
	}

	return user.ID(id), nil
}

func UserCreate(req *http.Request) (*user.User, error) {
	user, err := parseUserValue(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserCreate: %w", err)
	}
	return user, nil
}

func UserRead(req *http.Request) (user.ID, error) {
	userID, err := parseUserID(req)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.UserRead: %w", err)
	}
	return userID, nil
}

func UserUpdate(req *http.Request) (*user.User, error) {
	userID, err := parseUserID(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserUpdate: %w", err)
	}

	user, err := parseUserValue(req)
	if err != nil {
		return nil, fmt.Errorf("http-handle/request.UserUpdate: %w", err)
	}

	user.ID = userID
	return user, nil
}

func UserDelete(req *http.Request) (user.ID, error) {
	userID, err := parseUserID(req)
	if err != nil {
		return 0, fmt.Errorf("http-handle/request.UserDelete: %w", err)
	}
	return userID, nil
}
