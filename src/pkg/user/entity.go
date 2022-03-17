package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordMinLength = 8
)

type ID int

type Password interface {
	fmt.Stringer
	Verify([]byte) bool
	Length() int
	Hash() []byte
}

// impl Password
type password struct {
	hash   []byte
	length int
}

func NewPassword(plain []byte) (Password, error) {
	bin, err := bcrypt.GenerateFromPassword(plain, 10)
	if err != nil {
		return nil, fmt.Errorf("pkg/user.NewPassword: %w", err)
	}
	return &password{
		hash: bin,
		// TODO length Unicode
		length: len(plain),
	}, nil
}

func NewPasswordFromHash(hash []byte) Password {
	return &password{
		hash: hash,
		// NOTE
		// 長さは分からないため、hash は常に必要なパスワード長を満たしていると仮定する
		length: PasswordMinLength,
	}
}

func (password) String() string {
	// パスワードは常に伏せ字で表示する
	return "*****"
}

func (pw *password) Verify(plain []byte) bool {
	// TODO error handling
	return bcrypt.CompareHashAndPassword(pw.hash, plain) == nil
}

func (pw *password) Length() int {
	return pw.length
}

func (pw *password) Hash() []byte {
	return pw.hash
}

type User struct {
	ID       ID
	Name     string
	Password Password
}

func New(name string, pw Password) *User {
	return &User{
		Name:     name,
		Password: pw,
	}
}

func (u *User) Valid() bool {
	return len(u.Name) > 0 &&
		u.Password.Length() >= PasswordMinLength
}
