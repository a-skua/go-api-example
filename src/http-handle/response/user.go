package response

import (
	"api.example.com/pkg/user"
	"encoding/json"
	"fmt"
	"net/http"
)

// user response
func writeUser(w http.ResponseWriter, u *user.User) error {
	type User struct {
		ID       user.ID   `json:"id"`
		Name     user.Name `json:"name"`
		Password string    `json:"password"`
	}

	body := struct {
		User `json:"user"`
	}{
		User: User{
			ID:       u.ID,
			Name:     u.Name,
			Password: "*****",
		},
	}

	writeHeader(w)
	return json.NewEncoder(w).Encode(&body)
}

func UserCreate(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return fmt.Errorf("http-handle/reponse.UserCreate: %w", err)
	}

	return nil
}

func UserRead(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return fmt.Errorf("http-handle/reponse.UserRead: %w", err)
	}

	return nil
}

func UserUpdate(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return fmt.Errorf("http-handle/reponse.UserUpdate: %w", err)
	}

	return nil
}

func UserDelete(w http.ResponseWriter) error {
	body := struct {
		User struct{} `json:"user"`
	}{}

	writeHeader(w)
	err := json.NewEncoder(w).Encode(&body)
	if err != nil {
		return fmt.Errorf("http-handle/reponse.UserDelete: %w", err)
	}

	return nil
}
