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

// impl Server
type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(u *User) (*User, error) {
	ok := u.validCreate()
	if !ok {
		return nil, fmt.Errorf("user.Create: invalid user")
	}

	return s.repository.UserCreate(u)
}

func (s *server) Read(id ID) (*User, error) {
	ok := id.Valid()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Read: invalid user_id")
	}

	return s.repository.UserRead(id)
}

func (s *server) Update(u *User) (*User, error) {
	ok := u.validUpdate()
	if !ok {
		return nil, fmt.Errorf("pkg/user.Update: invalid user")
	}

	return s.repository.UserUpdate(u)
}

func (s *server) Delete(id ID) error {
	ok := id.Valid()
	if !ok {
		return fmt.Errorf("pkg/user.Delete: invalid user_id")
	}

	return s.repository.UserDelete(id)
}
