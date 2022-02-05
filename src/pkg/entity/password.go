package entity

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

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
