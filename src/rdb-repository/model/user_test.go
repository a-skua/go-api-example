package model

import (
	"api.example.com/env"
	users "api.example.com/pkg/user"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
	"time"
)

// mock
func newDB() *sql.DB {
	addr := env.Get("TEST_DB_ADDR")
	name := env.Get("TEST_DB_NAME")
	user := env.Get("TEST_DB_USER")
	password := env.GetSecure("TEST_DB_PASSWORD")

	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true",
		user.Value(),
		password.Value(),
		addr.Value(),
		name.Value(),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func TestNewUser(t *testing.T) {
	type test struct {
		testcase string
		in       *users.User
		want     User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := NewUser(tt.in)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := users.NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success",
				in: &users.User{
					ID:       1,
					Name:     "bob",
					Password: pw,
				},
				want: &user{
					ID:        1,
					Name:      "bob",
					Password:  pw.Hash(),
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserNewEntity(t *testing.T) {
	type test struct {
		testcase string
		user     User
		want     *users.User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.user.NewEntity()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			user: &user{
				ID:        2,
				Name:      "alice",
				Password:  []byte("dummy hash!"),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			want: &users.User{
				ID:       2,
				Name:     "alice",
				Password: users.NewPasswordFromHash([]byte("dummy hash!")),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserCreate(t *testing.T) {
	type test struct {
		testcase string
		user     *user
		db       *sql.DB
	}

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		tx, err := tt.db.Begin()
		if err != nil {
			panic(err)
		}
		defer tx.Rollback()

		beforeUser := *tt.user
		err = tt.user.Create(tx)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if reflect.DeepEqual(beforeUser, tt.user) {
			t.Fatalf("unchanged before=%v,  after=%v", beforeUser, tt.user)
		}

		rows, err := tx.Query("select * from users")
		if err != nil {
			panic(err)
		}

		want := tt.user
		got := &user{}

		if rows.Next() {
			err := rows.Scan(&got.ID, &got.Name, &got.Password, &got.CreatedAt, &got.UpdatedAt)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
		} else {
			t.Fatalf("failed create user")
		}

		{ // NOTE MySQL は micro sec まで?
			want.CreatedAt = want.CreatedAt.Round(time.Microsecond)
			want.UpdatedAt = want.UpdatedAt.Round(time.Microsecond)
			if want.CreatedAt.Equal(got.CreatedAt) && want.UpdatedAt.Equal(got.UpdatedAt) {
				got.CreatedAt = want.CreatedAt
				got.UpdatedAt = want.UpdatedAt
			} else {
				t.Fatalf("invalid time want=%v, got=%v.", want.CreatedAt, got.CreatedAt)
			}
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want=%v, got=%v.", want, got)
		}

		if rows.Next() {
			t.Fatalf("unexpected multiple users")
		}
	}

	tests := []test{
		{
			testcase: "success",
			user: &user{
				Name:     "bob",
				Password: []byte("password"),
			},
			db: newDB(),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		testcase string
		user     *user
		db       *sql.DB
	}

	do := func(tt test) {
		// TODO
	}

	tests := []test{}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type test struct {
		testcase string
		user     *user
		db       *sql.DB
		wantErr  bool
	}

	do := func(tt test) {
		// TODO
	}

	tests := []test{}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		testcase string
		user     *user
		db       *sql.DB
	}

	do := func(tt test) {
		// TODO
	}

	tests := []test{}

	for _, tt := range tests {
		do(tt)
	}
}
