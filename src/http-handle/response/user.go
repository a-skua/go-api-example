package response

import (
	"api.example.com/pkg/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, user *entity.User) error {
	type Company struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type User struct {
		ID        entity.UserID `json:"id"`
		Name      string        `json:"name"`
		Password  string        `json:"password"`
		Companies []Company     `json:"companeis"`
	}
	res := struct {
		User User `json:"user"`
	}{
		User: User{
			ID:        user.ID,
			Name:      user.Name,
			Password:  user.Password.String(),
			Companies: []Company{},
		},
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(&res)
	if err != nil {
		return fmt.Errorf("HTTP Response UserCreate: %w", err)
	}

	return nil
}
