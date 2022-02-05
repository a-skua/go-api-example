package entity

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// User ID
type UserID int

// Password
type Password interface {
	fmt.Stringer
	Verify(string) bool
	Length() int
	Hash() []byte
}

// implements Password
type password struct {
	hash   []byte
	length int
}

func NewPassword(plain string) (Password, error) {
	bin, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		return nil, fmt.Errorf("entity.NewPassword: %w", err)
	}
	return &password{
		hash:   bin,
		length: len(plain),
	}, nil
}

func PasswordFromHash(hash []byte) Password {
	return &password{
		hash:   hash,
		length: 8, // 仮置き
	}
}

func (pw *password) Verify(plain string) bool {
	return bcrypt.CompareHashAndPassword(pw.hash, []byte(plain)) == nil
}

func (pw *password) Length() int {
	return pw.length
}

func (password) String() string {
	return "*****"
}

func (pw *password) Hash() []byte {
	return pw.hash
}

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
