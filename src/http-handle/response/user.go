package response

import (
	"api.example.com/pkg/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

// user response
func user(w http.ResponseWriter, user *entity.User) error {
	type Company struct {
		ID   entity.CompanyID `json:"id"`
		Name string           `json:"name"`
	}
	type User struct {
		ID        entity.UserID `json:"id"`
		Name      string        `json:"name"`
		Password  string        `json:"password"`
		Companies []*Company    `json:"companies"`
	}

	companies := make([]*Company, 0, len(user.Companies))
	for _, c := range user.Companies {
		companies = append(companies, &Company{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	res := struct {
		User User `json:"user"`
	}{
		User: User{
			ID:        user.ID,
			Name:      user.Name,
			Password:  user.Password.String(),
			Companies: companies,
		},
	}

	writeHeader(w)
	return json.NewEncoder(w).Encode(&res)
}

func UserCreate(w http.ResponseWriter, u *entity.User) error {
	err := user(w, u)
	if err != nil {
		return fmt.Errorf("HTTP Response UserCreate: %w", err)
	}
	return nil
}

func UserRead(w http.ResponseWriter, u *entity.User) error {
	err := user(w, u)
	if err != nil {
		return fmt.Errorf("HTTP Response UserRead: %w", err)
	}
	return nil
}

func UserUpdate(w http.ResponseWriter, u *entity.User) error {
	err := user(w, u)
	if err != nil {
		return fmt.Errorf("HTTP Response UserUpdate: %w", err)
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
		return fmt.Errorf("HTTP Response UserDelete: %w", err)
	}

	return nil
}
