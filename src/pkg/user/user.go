package user

import (
	"fmt"
)

type ID int

func (id ID) Valid() bool {
	return id > 0
}

type Name string

// Name length ≥ 1
func (n Name) valid() bool {
	return len(n) >= 1
}

type PlainText = string

type Password interface {
	fmt.Stringer
	Verify(PlainText) bool
	Length() int
	Hash() []byte
}

const (
	PasswordMinLength = 8
	PasswordString    = "*****"
)

// Password length ≥ 8
func validPassword(p Password) bool {
	return p.Length() >= PasswordMinLength
}

type User struct {
	ID
	Name
	Password
}

func New(name Name, pw Password) *User {
	return &User{
		Name:     name,
		Password: pw,
	}
}

func (u *User) valid() bool {
	return u.Name.valid() && validPassword(u.Password)
}
