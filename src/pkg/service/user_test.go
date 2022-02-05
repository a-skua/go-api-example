package service

import (
	"api.example.com/pkg/entity"
	"errors"
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

// mock repository
type mockRepository struct {
	user *entity.User
	err  error
}

func (r *mockRepository) UserCreate(*entity.User) (*entity.User, error) {
	return r.user, r.err
}

func TestNewUser(t *testing.T) {
	r := &mockRepository{}
	want := &User{r}
	got := NewUser(r)
	if *want != *got {
		t.Fatalf("service.NewUser: want=%v, got=%v.", want, got)
	}
}

func TestUserCreate(t *testing.T) {
	// testing data
	tests := []struct {
		user *User
		in   *entity.User
		want *entity.User
	}{
		{
			&User{&mockRepository{user: &entity.User{
				ID:        1,
				Name:      "Bob",
				Password:  plainPassword("password"),
				Companies: nil,
			}}},
			&entity.User{
				Name:     "Bob",
				Password: plainPassword("password"),
			},
			&entity.User{
				ID:        1,
				Name:      "Bob",
				Password:  plainPassword("password"),
				Companies: nil,
			},
		},
	}

	// do tests
	for _, tt := range tests {
		got, err := tt.user.Create(tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("service.User.Create: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestFailedUserCreate(t *testing.T) {
	// testing data
	tests := []struct {
		user *User
		in   *entity.User
	}{
		{
			&User{&mockRepository{}},
			&entity.User{
				Name:     "Alice",
				Password: plainPassword("qwerty"),
			},
		},
		{
			&User{&mockRepository{
				err: errors.New("Repository Error"),
			}},
			&entity.User{
				Name:     "Alice",
				Password: plainPassword("passowrd"),
			},
		},
	}

	// do tests
	for _, tt := range tests {
		_, err := tt.user.Create(tt.in)
		if err == nil {
			t.Fatal("Expect Error")
		}
	}
}