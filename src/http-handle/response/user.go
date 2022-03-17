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
		ID       user.ID `json:"id"`
		Name     string  `json:"name"`
		Password string  `json:"password"`
	}

	res := struct {
		User User `json:"user"`
	}{
		User: User{
			ID:       u.ID,
			Name:     u.Name,
			Password: u.Password.String(),
		},
	}

	writeHeader(w)
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		return fmt.Errorf("http-handle/reponse.writeUser: %w", err)
	}
	return nil
}

func UserCreate(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return err
	}
	return nil
}

func UserRead(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return err
	}
	return nil
}

func UserUpdate(w http.ResponseWriter, u *user.User) error {
	err := writeUser(w, u)
	if err != nil {
		return err
	}
	return nil
}

func UserDelete(w http.ResponseWriter) error {
	res := struct {
		User struct{} `json:"user"`
	}{}

	writeHeader(w)
	enc := json.NewEncoder(w)
	err := enc.Encode(&res)
	if err != nil {
		return err
	}

	return nil
}
