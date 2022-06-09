package model

import (
	"api.example.com/env"
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
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

func testUserDateTime(want, got *user) error {
	if want.UpdatedAt.Equal(got.UpdatedAt) {
		want.UpdatedAt = got.UpdatedAt
	} else {
		return fmt.Errorf("UpdatedAt: want=%v, got=%v.", want.UpdatedAt, got.UpdatedAt)
	}

	if want.CreatedAt.Equal(got.CreatedAt) {
		want.CreatedAt = got.CreatedAt
	} else {
		return fmt.Errorf("CreatedAt: want=%v, got=%v.", want.CreatedAt, got.CreatedAt)
	}

	return nil
}

func TestNewUser(t *testing.T) {
	type test struct {
		testcase string
		in       *users.User
		want     User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := NewUser(tt.in)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				in: &users.User{
					ID:       1,
					Name:     "bob",
					Password: pw,
				},
				want: &user{
					ID:       1,
					Name:     "bob",
					Password: pw.Hash(),
				},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_NewEntity(t *testing.T) {
	type test struct {
		testcase string
		user     User
		want     *users.User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.user.NewEntity()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			user: &user{
				ID:       2,
				Name:     "alice",
				Password: []byte("password hash!"),
			},
			want: &users.User{
				ID:       2,
				Name:     "alice",
				Password: password.FromHash([]byte("password hash!")),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Create(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     *user
		want     *user
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()

			tx, err := tt.db.Begin()
			if err != nil {
				panic(err)
			}
			defer tx.Rollback()

			{
				now := currentTime() // 失敗する可能性あり
				err = tt.user.Create(tx)
				if tt.wantErr != (err != nil) {
					t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
				}
				tt.want.UpdatedAt = now
				tt.want.CreatedAt = now
				if !tt.user.ID.Valid() {
					t.Fatal("invalid user.ID")
				}
				tt.want.ID = tt.user.ID
			}
			if !reflect.DeepEqual(tt.want, tt.user) {
				t.Fatalf("want=%v, got=%v.", tt.want, tt.user)
			}

			rows, err := tx.Query("select id, name, password, updated_at, created_at from users")
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			got := &user{}
			if rows.Next() {
				err := rows.Scan(&got.ID, &got.Name, &got.Password, &got.UpdatedAt, &got.CreatedAt)
				if err != nil {
					panic(err)
				}
			} else {
				t.Fatal("failed create user")
			}

			err = testUserDateTime(tt.want, got)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			db: newDB(),
			user: &user{
				Name:     "bob",
				Password: []byte("password"),
			},
			want: &user{
				Name:     "bob",
				Password: []byte("password"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Read(t *testing.T) {
	type table struct {
		users []*user
	}

	type test struct {
		testcase string
		db       *sql.DB
		table
		want    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()

			tx, err := tt.db.Begin()
			if err != nil {
				panic(err)
			}
			defer tx.Rollback()
			for _, u := range tt.users {
				err := u.Create(tx)
				if err != nil {
					panic(err)
				}
			}

			got := &user{ID: tt.users[0].ID}
			tt.want.ID = tt.users[0].ID
			tt.want.UpdatedAt = tt.users[0].UpdatedAt
			tt.want.CreatedAt = tt.users[0].CreatedAt

			err = got.Read(tx)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			err = testUserDateTime(tt.want, got)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			db: newDB(),
			table: table{
				users: []*user{
					{
						Name:     "foo",
						Password: []byte("password"),
					},
					{
						Name:     "bar",
						Password: []byte("qwerty"),
					},
					{
						Name:     "baz",
						Password: []byte("12345678"),
					},
				},
			},
			want: &user{
				Name:     "foo",
				Password: []byte("password"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Update(t *testing.T) {
	type table struct {
		users []*user
	}

	type test struct {
		testcase string
		db       *sql.DB
		table
		user    *user
		want    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()

			tx, err := tt.db.Begin()
			if err != nil {
				panic(err)
			}
			defer tx.Rollback()

			for _, u := range tt.users {
				err := u.Create(tx)
				if err != nil {
					panic(err)
				}
			}

			time.Sleep(time.Second)

			{
				tt.user.ID = tt.users[0].ID
				now := currentTime()
				err := tt.user.Update(tx)
				if tt.wantErr != (err != nil) {
					t.Fatalf("want-error%v, error=%v.", tt.wantErr, err)
				}
				tt.want.ID = tt.users[0].ID
				tt.want.UpdatedAt = now
			}
			if !reflect.DeepEqual(tt.want, tt.user) {
				t.Fatalf("wnat=%v, got=%v.", tt.want, tt.user)
			}
			tt.want.CreatedAt = tt.users[0].CreatedAt

			got := &user{ID: tt.want.ID}
			err = got.Read(tx)
			if err != nil {
				panic(err)
			}

			err = testUserDateTime(tt.want, got)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			db: newDB(),
			table: table{
				users: []*user{
					{
						Name:     "foo",
						Password: []byte("password"),
					},
					{
						Name:     "bar",
						Password: []byte("qwerty"),
					},
					{
						Name:     "baz",
						Password: []byte("12345678"),
					},
				},
			},
			user: &user{
				Name:     "update foo!",
				Password: []byte("password!!"),
			},
			want: &user{
				Name:     "update foo!",
				Password: []byte("password!!"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Delete(t *testing.T) {
	type table struct {
		users []*user
	}

	type test struct {
		testcase string
		db       *sql.DB
		table
		user    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()

			tx, err := tt.db.Begin()
			if err != nil {
				panic(err)
			}
			defer tx.Rollback()

			for _, u := range tt.users {
				err := u.Create(tx)
				if err != nil {
					panic(err)
				}
			}

			tt.user.ID = tt.users[0].ID
			err = tt.user.Delete(tx)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			got := user{ID: tt.user.ID}
			err = got.Read(tx)
			if err == nil {
				t.Fatal("error is nil")
			}
			t.Log(err)
		})
	}

	tests := []*test{
		{
			db: newDB(),
			table: table{
				users: []*user{
					{
						Name:     "foo",
						Password: []byte("password"),
					},
					{
						Name:     "bar",
						Password: []byte("qwerty"),
					},
					{
						Name:     "baz",
						Password: []byte("12345678"),
					},
				},
			},
			user:    &user{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
