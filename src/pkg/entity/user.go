package entity

import (
	"fmt"
)

// User ID
type UserID int

// User Entity
type User struct {
	ID        UserID
	Name      string
	Password  Password
	Companies []*Company
}

func NewUser(name string, pw Password) *User {
	return &User{
		Name:     name,
		Password: pw,
	}
}

func (u *User) Validate() error {
	if l := len(u.Name); l == 0 {
		return fmt.Errorf("entity.User.Validate: Name.Length %d ≥ 1", l)
	}

	if l := u.Password.Length(); l < 8 {
		return fmt.Errorf("entity.User.Validate: Password.Length %d ≥ 8", l)
	}

	return nil
}
