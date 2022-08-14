package model

import (
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
	"time"
)

func TestPasswordHash_String(t *testing.T) {
	type test struct {
		name string
		hash passwordHash
		want string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hash.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			hash: passwordHash("password"),
			want: "password",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNewUser(t *testing.T) {
	type test struct {
		name string
		user *users.User
		want User
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.user)
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
				name: "true",
				user: &users.User{
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

func TestNewUserFromID(t *testing.T) {
	type test struct {
		name string
		id   users.ID
		want User
	}

	do := func(tt *test) {
		got := NewUserFromID(tt.id)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []*test{
		{
			name: "true",
			id:   1,
			want: &user{
				ID: 1,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_NewEntity(t *testing.T) {
	type test struct {
		name string
		user User
		want *users.User
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.user.NewEntity()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			user: &user{
				ID:        2,
				Name:      "alice",
				Password:  []byte("password hash!"),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			want: &users.User{
				ID:        2,
				Name:      "alice",
				Password:  password.FromHash([]byte("password hash!")),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Create(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		user    *users.User
		want    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.user).(*user)

			err := got.Create(tt.db)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()

			tt.want.ID = got.ID
			testDiffTime(t, wantTime, got.CreatedAt)
			tt.want.CreatedAt = got.CreatedAt
			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}

			// 作成されていることの確認
			err = db.
				QueryRow("select id, name, password, updated_at, created_at from users where id = ?", got.ID).
				Scan(&got.ID, &got.Name, &got.Password, &got.UpdatedAt, &got.CreatedAt)
			if err != nil {
				t.Fatal(err)
			}

			testDiffTime(t, wantTime, got.CreatedAt)
			tt.want.CreatedAt = got.CreatedAt
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
				name: "true",
				db:   db,
				user: users.New("bob", pw),
				want: &user{
					Name:     "bob",
					Password: pw.Hash(),
				},
				wantErr: false,
			}
		}(),
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "failed ExecContext",
				db: &testdb{
					err:         errors.New("test error"),
					execContext: true,
				},
				user:    users.New("alice", pw),
				want:    nil,
				wantErr: true,
			}
		}(),
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "failed result.LastInsertId",
				db: &testdb{
					result: &queryResult{
						err:          errors.New("test error"),
						lastInsertID: true,
					},
					execContext: true,
				},
				user:    users.New("alice", pw),
				want:    nil,
				wantErr: true,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Read(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		id      users.ID
		want    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserFromID(tt.id).(*user)
			err := got.Read(tt.db)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			testDiffTime(t, tt.want.CreatedAt, got.CreatedAt)
			tt.want.CreatedAt = got.CreatedAt
			testDiffTime(t, tt.want.UpdatedAt, got.UpdatedAt)
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
			model := NewUser(users.New("Bob", pw)).(*user)
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			return &test{
				name:    "true",
				db:      db,
				id:      model.ID,
				want:    model,
				wantErr: false,
			}
		}(),
		{
			name:    "failed scan",
			db:      db,
			id:      0,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_Update(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from users")

	type test struct {
		name    string
		db      DB
		user    *users.User
		want    *user
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.user).(*user)
			err := got.Update(tt.db)
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

			// 更新されていることの確認
			err = db.
				QueryRow("select id, name, password, updated_at, created_at from users where id = ?", got.ID).
				Scan(&got.ID, &got.Name, &got.Password, &got.UpdatedAt, &got.CreatedAt)
			if err != nil {
				t.Fatal(err)
			}

			testDiffTime(t, wantTime, got.UpdatedAt)
			tt.want.UpdatedAt = got.UpdatedAt
			testDiffTime(t, wantTime, got.CreatedAt)
			tt.want.CreatedAt = got.CreatedAt

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
			model := NewUser(users.New("Bob", pw)).(*user)
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "true",
				db:   db,
				user: &users.User{
					ID:       model.ID,
					Name:     "Alice",
					Password: newPW,
				},
				want: &user{
					ID:       model.ID,
					Name:     "Alice",
					Password: newPW.Hash(),
				},
				wantErr: false,
			}
		}(),
		func() *test {
			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "failed ExecContext",
				db: &testdb{
					err:         errors.New("test error"),
					execContext: true,
				},
				user: &users.User{
					ID:       1,
					Name:     "Alice",
					Password: newPW,
				},
				want:    nil,
				wantErr: true,
			}
		}(),
		func() *test {
			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "failed result.RowsAffected",
				db: &testdb{
					result: &queryResult{
						err:          errors.New("test error"),
						rowsAffected: true,
					},
					execContext: true,
				},
				user: &users.User{
					ID:       1,
					Name:     "Alice",
					Password: newPW,
				},
				want:    nil,
				wantErr: true,
			}
		}(),
		func() *test {
			newPW, err := password.New("12345678")
			if err != nil {
				panic(err)
			}

			return &test{
				name: "invalid rows-affected",
				db: &testdb{
					result: &queryResult{
						rows:         0,
						rowsAffected: true,
					},
					execContext: true,
				},
				user: &users.User{
					ID:       1,
					Name:     "Alice",
					Password: newPW,
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

func TestUser_Delete(t *testing.T) {
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
			got := NewUserFromID(tt.id).(*user)
			err := got.Delete(tt.db)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			var count int64
			err = db.
				QueryRow("select count(*) from users where id = ?", tt.id).
				Scan(&count)
			if count != 0 {
				t.Fatalf("failed delete id=%v, count=%v.", tt.id, count)
			}
		})
	}

	tests := []*test{
		func() *test {
			pw, err := password.New("password")
			if err != nil {
				panic(err)
			}
			model := NewUser(users.New("Bob", pw)).(*user)
			err = model.Create(db)
			if err != nil {
				panic(err)
			}

			return &test{
				name:    "true",
				db:      db,
				id:      model.ID,
				wantErr: false,
			}
		}(),
		{
			name: "failed ExecContext",
			db: &testdb{
				err:         errors.New("test error"),
				execContext: true,
			},
			id:      1,
			wantErr: true,
		},
		{
			name: "failed result.RowsAffected",
			db: &testdb{
				result: &queryResult{
					err:          errors.New("test error"),
					rowsAffected: true,
				},
				execContext: true,
			},
			id:      1,
			wantErr: true,
		},
		{
			name: "invalid rows-affected",
			db: &testdb{
				result: &queryResult{
					rows:         0,
					rowsAffected: true,
				},
				execContext: true,
			},
			id:      1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
