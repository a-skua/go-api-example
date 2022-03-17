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

	do := func(tt test) {
		defer tt.db.Close()

		t.Logf("testcase: %s", tt.testcase)

		got := func() *rdb {
			tmp := New(tt.db)
			got, ok := tmp.(*rdb)
			if !ok {
				t.Fatalf("type want=%T, got=%T.", got, tmp)
			}
			return got
		}()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			db := newDB()

			return test{
				testcase: "success",
				db:       db,
				want: &rdb{
					db: db,
				},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepositoryClose(t *testing.T) {
	type test struct {
		testcase   string
		repository Repository
		wantErr    bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		err := tt.repository.Close()
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}
	}

	tests := []test{
		func() test {
			db := newDB()

			return test{
				testcase:   "success 1",
				repository: &rdb{db: db},
				wantErr:    false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}
