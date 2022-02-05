package entity

import (
	"reflect"
	"testing"
)

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
