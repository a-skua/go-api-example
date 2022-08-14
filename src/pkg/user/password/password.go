package password

import (
	"api.example.com/pkg/user"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	generatedPassword = bcrypt.GenerateFromPassword
	comparePassword   = bcrypt.CompareHashAndPassword
)

// impl Password
type password struct {
	hash   []byte
	length int
}

func New(plain string) (user.Password, error) {
	bin, err := generatedPassword([]byte(plain), 10)
	if err != nil {
		return nil, fmt.Errorf("user.NewPassword: %w", err)
	}
	return &password{
		hash:   bin,
		length: len(plain),
	}, nil
}

func FromHash(hash []byte) user.Password {
	return &password{
		hash:   hash,
		length: 0,
	}
}

func (pw *password) Verify(plain string) bool {
	return comparePassword(pw.hash, []byte(plain)) == nil
}

func (pw *password) Length() int {
	return pw.length
}

func (pw *password) Hash() []byte {
	return pw.hash
}
