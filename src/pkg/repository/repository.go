package repository

import (
	"api.example.com/pkg/entity"
)

type User interface {
	UserCreate(*entity.User) (*entity.User, error)
	UserRead(entity.UserID) (*entity.User, error)
}
