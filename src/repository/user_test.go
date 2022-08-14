package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"api.example.com/repository/model"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type user struct {
	entity *users.User
	err    error
	// flags
	create, read, update, delete bool
}

func (u *user) Create(tx model.DB) error {
	if u.create {
		return u.err
	}
	return fmt.Errorf("invalid create")
}

func (u *user) Read(tx model.DB) error {
	if u.read {
		return u.err
	}
	return fmt.Errorf("invalid read")
}

func (u *user) Update(tx model.DB) error {
	if u.update {
		return u.err
	}
	return fmt.Errorf("invald update")
}

func (u *user) Delete(tx model.DB) error {
	if u.delete {
		return u.err
	}
	return fmt.Errorf("invalid delete")
}

func (u *user) NewEntity() *users.User {
	return u.entity
}

func TestUserCreate(t *testing.T) {
	type test struct {
		name    string
		tx      Transaction
		user    model.User
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserCreate(tt.tx, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			tx: &transaction{
				commit: true,
			},
			user: &user{
				entity: &users.User{
					ID:        1,
					Name:      "bob",
					Password:  password.FromHash([]byte("password")),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			},
			want: &users.User{
				ID:        1,
				Name:      "bob",
				Password:  password.FromHash([]byte("password")),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "failed create",
			tx: &transaction{
				rollback: true,
			},
			user: &user{
				err:    errors.New("test error"),
				create: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed commit",
			tx: &transaction{
				errCommit: errors.New("test error"),
				commit:    true,
			},
			user: &user{
				entity: &users.User{
					ID:        1,
					Name:      "bob",
					Password:  password.FromHash([]byte("password")),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		name    string
		db      DB
		user    model.User
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserRead(tt.db, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			db:   &mockDB{},
			user: &user{
				entity: &users.User{
					ID:        1,
					Name:      "bob",
					Password:  password.FromHash([]byte("password")),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				read: true,
			},
			want: &users.User{
				ID:        1,
				Name:      "bob",
				Password:  password.FromHash([]byte("password")),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "failed read",
			db:   &mockDB{},
			user: &user{
				err:  errors.New("test error"),
				read: true,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type test struct {
		name    string
		tx      Transaction
		user    model.User
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserUpdate(tt.tx, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			tx: &transaction{
				commit: true,
			},
			user: &user{
				entity: &users.User{
					ID:        1,
					Name:      "bob",
					Password:  password.FromHash([]byte("password")),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			},
			want: &users.User{
				ID:        1,
				Name:      "bob",
				Password:  password.FromHash([]byte("password")),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "failed update",
			tx: &transaction{
				rollback: true,
			},
			user: &user{
				err:    errors.New("test error"),
				update: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed commit",
			tx: &transaction{
				errCommit: errors.New("test error"),
				commit:    true,
			},
			user: &user{
				entity: &users.User{
					ID:        1,
					Name:      "bob",
					Password:  password.FromHash([]byte("password")),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		name    string
		tx      Transaction
		user    model.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDelete(tt.tx, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			tx: &transaction{
				commit: true,
			},
			user: &user{
				delete: true,
			},
			wantErr: false,
		},
		{
			name: "failed delete",
			tx: &transaction{
				commit: true,
			},
			user: &user{
				err:    errors.New("test error"),
				delete: true,
			},
			wantErr: true,
		},
		{
			name: "failed commit",
			tx: &transaction{
				errCommit: errors.New("test error"),
				commit:    true,
			},
			user: &user{
				delete: true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
