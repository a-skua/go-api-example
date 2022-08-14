package user

import (
	"time"
)

type ID int

func (id ID) valid() bool {
	return id > 0
}

type Name string

// 1 ≤ name.length ≤ 255
func (n Name) valid() bool {
	return len(n) > 0 && len(n) < 256
}

type PlainPassword = string

type Password interface {
	Verify(PlainPassword) bool
	Length() int
	Hash() []byte
}

// 8 ≤ password.length ≤ 255
func validPassword(p Password) bool {
	return p.Length() > 7 && p.Length() < 256
}

type User struct {
	ID        ID
	Name      Name
	Password  Password
	UpdatedAt time.Time
}

func New(name Name, pw Password) *User {
	return &User{
		Name:     name,
		Password: pw,
	}
}

// valid user when create
func (u *User) validCreate() bool {
	return u.Name.valid() && validPassword(u.Password)
}

// vlid user when update
func (u *User) validUpdate() bool {
	return u.ID.valid() && u.validCreate()
}
