package entity

import (
	"reflect"
	"testing"
)

// mock Password
type plainPassword string

func (pw plainPassword) Verify(plain string) bool {
	return pw == plainPassword(plain)
}

func (pw plainPassword) Length() int {
	return len(pw)
}

func (plainPassword) String() string {
	return "*****"
}

func TestNewPassword(t *testing.T) {
	_, err := NewPassword("password")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPasswordVerify(t *testing.T) {
	pw1, _ := NewPassword("password")
	pw2, _ := NewPassword("password")

	tests := []struct {
		password *password
		in       string
		want     bool
	}{
		{
			pw1.(*password),
			"password",
			true,
		},
		{
			pw2.(*password),
			"passworld",
			false,
		},
	}

	for _, tt := range tests {
		got := tt.password.Verify(tt.in)
		if tt.want != got {
			t.Fatalf("entity.password.Verify: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestPasswordLength(t *testing.T) {
	pw1, _ := NewPassword("password")
	pw2, _ := NewPassword("qwerty")

	tests := []struct {
		password *password
		want     int
	}{
		{
			pw1.(*password),
			8,
		},
		{
			pw2.(*password),
			6,
		},
	}

	for _, tt := range tests {
		got := tt.password.Length()
		if tt.want != got {
			t.Fatalf("entity.password.Length: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestPasswordString(t *testing.T) {
	pw1, _ := NewPassword("password")
	pw2, _ := NewPassword("qwerty")

	tests := []struct {
		password *password
		want     string
	}{
		{
			pw1.(*password),
			"*****",
		},
		{
			pw2.(*password),
			"*****",
		},
	}

	for _, tt := range tests {
		got := tt.password.String()
		if tt.want != got {
			t.Fatalf("entity.password.Length: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestNewUser(t *testing.T) {
	want := &User{Name: "Bob", Password: plainPassword("password")}
	got := NewUser("Bob", plainPassword("password"))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("entity.NewUser: want=%v, got=%v.", want, got)
	}
}

func TestUserValidate(t *testing.T) {
	pw, _ := NewPassword("password")

	tests := []struct {
		user *User
	}{
		{&User{Name: "Bob", Password: pw}},
	}

	for _, tt := range tests {
		err := tt.user.Validate()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestFailedUserValidate(t *testing.T) {
	pw1, _ := NewPassword("password")
	pw2, _ := NewPassword("qwerty")

	tests := []struct {
		user *User
	}{
		{&User{Name: "", Password: pw1}},
		{&User{Name: "Alice", Password: pw2}},
	}

	for _, tt := range tests {
		err := tt.user.Validate()
		if err == nil {
			t.Fatal("Expect Error")
		}
	}
}
