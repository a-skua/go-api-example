package user

import (
	"fmt"
)

type Service interface {
	Create(*User) (*User, error)
	Read(ID) (*User, error)
	Update(*User) (*User, error)
	Delete(ID) error
}

// impl Service
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(u *User) (*User, error) {
	ok := u.Valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.service.Create: invalid user")
	}

	u, err := s.repository.UserCreate(u)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.service.Create: %w", err)
	}

	return u, nil
}

func (s *service) Read(id ID) (*User, error) {
	u, err := s.repository.UserRead(id)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.service.Read: %w", err)
	}

	return u, nil
}

func (s *service) Update(u *User) (*User, error) {
	ok := u.Valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.service.Update: invalid user")
	}

	u, err := s.repository.UserUpdate(u)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.service.Update: %w", err)
	}

	return u, nil
}

func (s *service) Delete(id ID) error {
	err := s.repository.UserDelete(id)
	if err != nil {
		return fmt.Errorf("pkg/user.service.Delete: %w", err)
	}

	return nil
}
