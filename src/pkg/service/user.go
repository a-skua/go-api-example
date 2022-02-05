package service

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
)

type User struct {
	repository repository.User
}

func NewUser(r repository.User) *User {
	return &User{r}
}

func (u *User) Create(newUser *entity.User) (*entity.User, error) {
	err := newUser.Validate()
	if err != nil {
		return nil, err
	}

	return u.repository.UserCreate(newUser)
}
