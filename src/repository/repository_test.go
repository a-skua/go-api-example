package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sync"
	"testing"
	"time"

	"api.example.com/env"
	companies "api.example.com/pkg/company"
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"api.example.com/repository/model"
	_ "github.com/go-sql-driver/mysql"
)

var tableLock sync.Mutex

func testDiffTime(t *testing.T, want, got time.Time) {
	t.Helper()

	if math.Abs(float64(want.Sub(got))) > float64(time.Second) {
		// 1秒以上ずれる場合はエラーとする
		t.Fatalf("updated_at want=%v, got=%v.", want, got)
	}
}

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

// mock
type mockDB struct {
	tx  *sql.Tx
	err error
	// flag
	begin, close bool
}

func (db *mockDB) Begin() (*sql.Tx, error) {
	if db.begin {
		return db.tx, db.err
	}
	return nil, errors.New("invalid Begin")
}

func (db *mockDB) Close() error {
	if db.close {
		return db.err
	}
	return errors.New("invalid Close")
}

func (db *mockDB) ExecContext(context.Context, string, ...interface{}) (sql.Result, error) {
	panic("invalid ExecContext")
}

func (db *mockDB) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	panic("invalid QueryRowContext")
}

type transaction struct {
	errCommit   error
	errRollback error
	// flag
	commit, rollback bool
}

func (tx *transaction) Commit() error {
	if tx.commit {
		return tx.errCommit
	}
	return errors.New("invalid Commit")
}

func (tx *transaction) Rollback() error {
	if tx.rollback {
		return tx.errRollback
	}
	return errors.New("invalid Rollback")
}

func (tx *transaction) ExecContext(context.Context, string, ...interface{}) (sql.Result, error) {
	panic("invalid ExecContext")
}

func (tx *transaction) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	panic("invalid QueryRowContext")
}

// test
func TestNew(t *testing.T) {
	type test struct {
		name string
		db   *sql.DB
		want Repository
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
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
				want: &repository{db},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_Close(t *testing.T) {
	type test struct {
		name       string
		repository Repository
		wantErr    bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.repository.Close()
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			name:       "true",
			repository: &repository{newDB()},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_UserCreate(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		user    *users.User
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{tt.db}
			got, err := repo.UserCreate(tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			tt.want.ID = got.ID
			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

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
				name:    "true",
				db:      db,
				user:    users.New("Bob", pw),
				want:    users.New("Bob", password.FromHash(pw.Hash())),
				wantErr: false,
			}
		}(),
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "failed begin-transaction",
				db: &mockDB{
					err:   errors.New("test error"),
					begin: true,
				},
				user:    users.New("Bob", pw),
				want:    users.New("Bob", password.FromHash(pw.Hash())),
				wantErr: true,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_UserRead(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	repo := New(db).(*repository)
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		id      users.ID
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.UserRead(tt.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

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
			model := model.NewUser(users.New("Bob", pw))
			err = model.Create(repo.db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()

			return &test{
				name: "true",
				id:   entity.ID,
				want: &users.User{
					ID:       entity.ID,
					Name:     "Bob",
					Password: password.FromHash(pw.Hash()),
				},
				wantErr: false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_UserUpdate(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		user    *users.User
		want    *users.User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{tt.db}
			got, err := repo.UserUpdate(tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			tt.want.ID = got.ID
			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

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

			model := model.NewUser(users.New("Bob", pw))
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()
			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "true",
				db:   db,
				user: &users.User{
					ID:       entity.ID,
					Name:     "Alice",
					Password: password.FromHash(newPW.Hash()),
				},
				want: &users.User{
					ID:       entity.ID,
					Name:     "Alice",
					Password: password.FromHash(newPW.Hash()),
				},
				wantErr: false,
			}
		}(),
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}

			model := model.NewUser(users.New("Hoge", pw))
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()
			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "faile begin-transaction",
				db: &mockDB{
					err:   errors.New("test error"),
					begin: true,
				},
				user: &users.User{
					ID:       entity.ID,
					Name:     "Alice",
					Password: password.FromHash(newPW.Hash()),
				},
				want:    nil,
				wantErr: true,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_UserDelete(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		id      users.ID
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{tt.db}
			err := repo.UserDelete(tt.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}
			model := model.NewUser(users.New("Bob", pw))
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()

			return &test{
				name:    "true",
				db:      db,
				id:      entity.ID,
				wantErr: false,
			}
		}(),
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}
			model := model.NewUser(users.New("Alice", pw))
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()

			return &test{
				name: "failed begin-transaction",
				db: &mockDB{
					err:   errors.New("test error"),
					begin: true,
				},
				id:      entity.ID,
				wantErr: true,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_CompanyCreate(t *testing.T) {
	t.Skip("TODO OwnerID")

	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from companies")

	type test struct {
		name    string
		db      DB
		company *companies.Company
		want    *companies.Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&repository{tt.db}).CompanyCreate(tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			tt.want.ID = got.ID
			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:    "ok",
			db:      db,
			company: companies.New("GREATE COMPANY", 1),
			want:    companies.New("GREATE COMPANY", 0), // TODO
			wantErr: false,
		},
		{
			name: "failed begin-transaction",
			db: &mockDB{
				err:   errors.New("test error"),
				begin: true,
			},
			company: companies.New("GREATE COMPANY", 1),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestRepository_CompanyRead(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	repo := New(db).(*repository)
	defer db.Close()
	defer db.Exec("delete from companies")

	type test struct {
		name    string
		id      companies.ID
		want    *companies.Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CompanyRead(tt.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		func() *test {
			model := model.NewCompany(companies.New("testCompany", 1))
			err := model.Create(repo.db)
			if err != nil {
				panic(err)
			}

			entity := model.NewEntity()

			return &test{
				name: "true",
				id:   entity.ID,
				want: &companies.Company{
					ID:      entity.ID,
					Name:    "testCompany",
					OwnerID: entity.OwnerID,
				},
				wantErr: false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}
