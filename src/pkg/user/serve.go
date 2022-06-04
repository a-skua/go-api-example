package user

import (
	"fmt"
)

type Repository interface {
	UserCreate(*User) (*User, error)
	UserRead(ID) (*User, error)
	UserUpdate(*User) (*User, error)
	UserDelete(ID) error
}

type Server interface {
	Create(*User) (*User, error)
	Read(ID) (*User, error)
	Update(*User) (*User, error)
	Delete(ID) error
}

// impl
type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(u *User) (*User, error) {
	ok := u.Valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Create: invalid user")
	}

	u, err := s.repository.UserCreate(u)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.Create: %w", err)
	}

	return u, nil
}

func (s *server) Read(id ID) (*User, error) {
	ok := id.valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Read: invalid user_id")
	}

	u, err := s.repository.UserRead(id)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.Read: %w", err)
	}

	return u, nil
}

func (s *server) Update(u *User) (*User, error) {
	ok := u.Valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Update: invalid user")
	}

	ok = u.ID.valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Update: invalid user_id")
	}

	u, err := s.repository.UserUpdate(u)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.Update: %w", err)
	}

	return u, nil
}

func (s *server) Delete(id ID) error {
	ok := id.valid()
	if !ok {
		return fmt.Errorf("pkg/user.Delete: invalid user_id")
	}

	err := s.repository.UserDelete(id)
	if err != nil {
		return fmt.Errorf("pkg/user.Delete: %w", err)
	}

	return nil
}
