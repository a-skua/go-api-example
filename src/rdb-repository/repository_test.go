package repository

import (
	"api.example.com/env"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
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

// test
func TestNew(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		want     Repository
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()
			got := New(tt.db)

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})

	}

	tests := []*test{
		func() *test {
			db := newDB()

			return &test{
				db:   db,
				want: &rdb{db},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRDB_Close(t *testing.T) {
	type test struct {
		testcase   string
		repository Repository
		wantErr    bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			err := tt.repository.Close()
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			repository: &rdb{newDB()},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
