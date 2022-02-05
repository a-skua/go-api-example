package service

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"fmt"
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

func (u *User) Read(userID, authID entity.UserID) (*entity.User, error) {
	if userID != authID {
		return nil, fmt.Errorf("service.User.Read: Unauthorized")
	}

	return u.repository.UserRead(userID)
}

func (u *User) Update(user *entity.User, authID entity.UserID) (*entity.User, error) {
	if user.ID != authID {
		return nil, fmt.Errorf("service.User.Update: Unauthorized")
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return u.repository.UserUpdate(user)
}

func (u *User) Delete(userID, authID entity.UserID) error {
	if userID != authID {
		return fmt.Errorf("service.User.Delete: Unauthorized")
	}

	return u.repository.UserDelete(userID)
}
