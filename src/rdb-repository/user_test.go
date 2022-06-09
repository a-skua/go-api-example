package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
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
	} else {
		return fmt.Errorf("invalid create")
	}
}

func (u *user) Read(tx model.DB) error {
	if u.read {
		return u.err
	} else {
		return fmt.Errorf("invalid read")
	}
}

func (u *user) Update(tx model.DB) error {
	if u.update {
		return u.err
	} else {
		return fmt.Errorf("invald update")
	}
}

func (u *user) Delete(tx model.DB) error {
	if u.delete {
		return u.err
	} else {
		return fmt.Errorf("invalid delete")
	}
}

func (u *user) NewEntity() *users.User {
	return u.entity
}

// test
func Test_userCreate(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		want     model.User
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

			got, err := userCreate(tx, tt.user)
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
			db:      newDB(),
			user:    &user{create: true},
			want:    &user{create: true},
			wantErr: false,
		},
		{
			db: newDB(),
			user: &user{
				create: true,
				err:    fmt.Errorf("failed create"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func Test_userRead(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		want     model.User
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			defer tt.db.Close()

			got, err := userRead(tt.db, tt.user)
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
			db:      newDB(),
			user:    &user{read: true},
			want:    &user{read: true},
			wantErr: false,
		},
		{
			db: newDB(),
			user: &user{
				read: true,
				err:  fmt.Errorf("failed read"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func Test_userUpdate(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		want     model.User
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

			got, err := userUpdate(tx, tt.user)
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
			db:      newDB(),
			user:    &user{update: true},
			want:    &user{update: true},
			wantErr: false,
		},
		{
			db: newDB(),
			user: &user{
				update: true,
				err:    fmt.Errorf("failed update"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func Test_userDelete(t *testing.T) {
	type test struct {
		testcase string
		db       *sql.DB
		user     model.User
		want     model.User
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

			err = userDelete(tx, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			db:      newDB(),
			user:    &user{delete: true},
			wantErr: false,
		},
		{
			db: newDB(),
			user: &user{
				delete: true,
				err:    fmt.Errorf("failed delete"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserCreate(t *testing.T) {
	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		testcase string
		user     *users.User
		want     *users.User
		wantErr  bool
	}

	do := func(tt *test) {
		repo := New(db)
		got, err := repo.UserCreate(tt.user)
		if tt.wantErr != (err != nil) {
			t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
		}

		if got != nil {
			tt.want.ID = got.ID
		}
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []*test{
		{
			user:    users.New("foo", password.FromHash([]byte("bar"))),
			want:    users.New("foo", password.FromHash([]byte("bar"))),
			wantErr: false,
		},
		{
			user:    users.New("foo", password.FromHash([]byte("bar"))),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		testcase string
		id       users.ID
		want     *users.User
		wantErr  bool
	}

	do := func(tt *test) {
		repo := New(db)
		got, err := repo.UserRead(tt.id)
		if tt.wantErr != (err != nil) {
			t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	models := []model.User{
		model.NewUser(users.New("foo", password.FromHash([]byte("password")))),
		model.NewUser(users.New("bar", password.FromHash([]byte("password")))),
		model.NewUser(users.New("baz", password.FromHash([]byte("password")))),
	}
	for _, m := range models {
		err := m.Create(db)
		if err != nil {
			panic(err)
		}
	}

	tests := []*test{
		{
			id:      models[0].NewEntity().ID,
			want:    models[0].NewEntity(),
			wantErr: false,
		},
		{
			id:      0,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		testcase string
		user     *users.User
		want     *users.User
		wantErr  bool
	}

	do := func(tt *test) {
		repo := New(db)
		got, err := repo.UserUpdate(tt.user)
		if tt.wantErr != (err != nil) {
			t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	models := []model.User{
		model.NewUser(users.New("foo", password.FromHash([]byte("password")))),
		model.NewUser(users.New("bar", password.FromHash([]byte("password")))),
		model.NewUser(users.New("baz", password.FromHash([]byte("password")))),
	}
	for _, m := range models {
		err := m.Create(db)
		if err != nil {
			panic(err)
		}
	}

	tests := []*test{
		{
			user: &users.User{
				ID:       models[0].NewEntity().ID,
				Name:     "fool",
				Password: password.FromHash([]byte("qwerty")),
			},
			want: &users.User{
				ID:       models[0].NewEntity().ID,
				Name:     "fool",
				Password: password.FromHash([]byte("qwerty")),
			},
			wantErr: false,
		},
		{
			user: &users.User{
				ID:       models[0].NewEntity().ID,
				Name:     "bar",
				Password: password.FromHash([]byte("qwerty")),
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
	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		testcase string
		id       users.ID
		wantErr  bool
	}

	do := func(tt *test) {
		repo := New(db)
		err := repo.UserDelete(tt.id)
		if tt.wantErr != (err != nil) {
			t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
		}
	}

	models := []model.User{
		model.NewUser(users.New("foo", password.FromHash([]byte("password")))),
		model.NewUser(users.New("bar", password.FromHash([]byte("password")))),
		model.NewUser(users.New("baz", password.FromHash([]byte("password")))),
	}
	for _, m := range models {
		err := m.Create(db)
		if err != nil {
			panic(err)
		}
	}

	tests := []*test{
		{
			id:      models[0].NewEntity().ID,
			wantErr: false,
		},
		{
			id:      0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
