package repository

import (
	"api.example.com/pkg/entity"
)

type User interface {
	UserCreate(*entity.User) (*entity.User, error)
}
