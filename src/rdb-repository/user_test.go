package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

// mock
type userModel struct {
	user *users.User
	err  error
}

func (u *userModel) Create(model.DB) error {
	return u.err
}

func (u *userModel) Read(model.DB) error {
	return u.err
}

func (u *userModel) Update(model.DB) error {
	return u.err
}

func (u *userModel) Delete(model.DB) error {
	return u.err
}

func (u *userModel) NewEntity() *users.User {
	return u.user
}

// test
func TestUserCreate(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		wantErr  bool
		want     *users.User
	}

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		got, err := userCreate(tt.db, tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: nil,
			},
			wantErr: false,
			want: &users.User{
				ID:   1,
				Name: "bob",
			},
		},
		{
			testcase: "failed",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		wantErr  bool
		want     *users.User
	}

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		got, err := userRead(tt.db, tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: nil,
			},
			wantErr: false,
			want: &users.User{
				ID:   1,
				Name: "bob",
			},
		},
		{
			testcase: "failed",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		wantErr  bool
		want     *users.User
	}

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		got, err := userUpdate(tt.db, tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: nil,
			},
			wantErr: false,
			want: &users.User{
				ID:   1,
				Name: "bob",
			},
		},
		{
			testcase: "failed",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		wantErr  bool
	}

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		err := userDelete(tt.db, tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}
	}

	tests := []test{
		{
			testcase: "success",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			testcase: "failed",
			db:       newDB(),
			user: &userModel{
				user: &users.User{
					ID:   1,
					Name: "bob",
				},
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
