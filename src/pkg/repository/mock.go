package repository

import (
	"api.example.com/pkg/entity"
)

// テスト用 Mock
type Mock struct {
	User  *entity.User
	Error error
}

func (r *Mock) UserCreate(*entity.User) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserRead(entity.UserID) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserUpdate(*entity.User) (*entity.User, error) {
	return r.User, r.Error
}

func (r *Mock) UserDelete(entity.UserID) error {
	return r.Error
}
