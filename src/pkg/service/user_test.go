package service

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
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

func (pw plainPassword) Hash() []byte {
	return []byte(pw)
}

func TestNewUser(t *testing.T) {
	r := &repository.Mock{}
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
			&User{&repository.Mock{User: &entity.User{
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
			&User{&repository.Mock{}},
			&entity.User{
				Name:     "Alice",
				Password: plainPassword("qwerty"),
			},
		},
		{
			&User{&repository.Mock{Error: errors.New("Repository Error")}},
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

func TestUserRead(t *testing.T) {
	tests := []struct {
		user *User
		in   entity.UserID
		auth entity.UserID
		want *entity.User
	}{
		{
			&User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			1,
			1,
			&entity.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			},
		},
	}

	for _, tt := range tests {
		got, err := tt.user.Read(tt.in, tt.auth)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("service.User.Read: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestFailedUserRead(t *testing.T) {
	tests := []struct {
		user *User
		in   entity.UserID
		auth entity.UserID
	}{
		{
			&User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			1,
			0,
		},
		{
			&User{&repository.Mock{Error: errors.New("Repository Error")}},
			1,
			1,
		},
	}

	for _, tt := range tests {
		_, err := tt.user.Read(tt.in, tt.auth)
		if err == nil {
			t.Fatalf("Expect Error")
		}
		t.Log(err)
	}
}

func TestUserUpdate(t *testing.T) {
	tests := []struct {
		user *User
		in   *entity.User
		auth entity.UserID
		want *entity.User
	}{
		{
			user: &User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Alice",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			in: &entity.User{
				ID:        1,
				Name:      "Aclice",
				Password:  plainPassword("password"),
				Companies: []*entity.Company{},
			},
			auth: 1,
			want: &entity.User{
				ID:       1,
				Name:     "Alice",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			},
		},
	}

	for _, tt := range tests {
		got, err := tt.user.Update(tt.in, tt.auth)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("service.User.Read: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestFailedUserUpdate(t *testing.T) {
	tests := []struct {
		user *User
		in   *entity.User
		auth entity.UserID
	}{
		{
			user: &User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Alice",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			in: &entity.User{
				ID:        1,
				Name:      "Aclice",
				Password:  plainPassword("password"),
				Companies: []*entity.Company{},
			},
			auth: 2,
		},
		{
			user: &User{&repository.Mock{Error: errors.New("Repository Error")}},
			in: &entity.User{
				ID:        1,
				Name:      "Aclice",
				Password:  plainPassword("password"),
				Companies: []*entity.Company{},
			},
			auth: 1,
		},
		{
			user: &User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Alice",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			in: &entity.User{
				ID:        1,
				Name:      "Aclice",
				Password:  plainPassword("qwerty"),
				Companies: []*entity.Company{},
			},
			auth: 1,
		},
	}

	for _, tt := range tests {
		_, err := tt.user.Update(tt.in, tt.auth)
		if err == nil {
			t.Fatalf("Expect Error")
		}
		t.Log(err)
	}
}

func TestUserDelete(t *testing.T) {
	tests := []struct {
		user *User
		in   entity.UserID
		auth entity.UserID
	}{
		{
			user: &User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			in:   1,
			auth: 1,
		},
	}

	for _, tt := range tests {
		err := tt.user.Delete(tt.in, tt.auth)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestFailedUserDelete(t *testing.T) {
	tests := []struct {
		user *User
		in   entity.UserID
		auth entity.UserID
	}{
		{
			user: &User{&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}},
			in:   1,
			auth: 2,
		},
		{
			user: &User{&repository.Mock{Error: errors.New("Repository Error")}},
			in:   1,
			auth: 1,
		},
	}

	for _, tt := range tests {
		err := tt.user.Delete(tt.in, tt.auth)
		if err == nil {
			t.Fatalf("Expect Error")
		}
		t.Log(err)
	}
}
