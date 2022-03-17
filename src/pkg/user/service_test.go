package user

import (
	"fmt"
	"reflect"
	"testing"
)

// mock
type mockRepository struct {
	user *User
	err  error
}

func (r *mockRepository) UserCreate(*User) (*User, error) {
	return r.user, r.err
}

func (r *mockRepository) UserRead(ID) (*User, error) {
	return r.user, r.err
}

func (r *mockRepository) UserUpdate(*User) (*User, error) {
	return r.user, r.err
}

func (r *mockRepository) UserDelete(ID) error {
	return r.err
}

// test
func TestNewService(t *testing.T) {
	type test struct {
		testcase   string
		repository Repository
		want       Service
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := NewService(tt.repository)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase:   "case 1",
			repository: &mockRepository{},
			want: &service{
				repository: &mockRepository{},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServiceCreate(t *testing.T) {
	type test struct {
		testcase string
		service  Service
		user     *User
		wantErr  bool
		want     *User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got, err := tt.service.Create(tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 1",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				user:    &User{Name: "bob", Password: pw},
				wantErr: false,
				want:    &User{ID: 1, Name: "bob", Password: pw},
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error invalid user",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				user:    &User{Name: "", Password: pw},
				wantErr: true,
				want:    nil,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error repository",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  fmt.Errorf("test error"),
				}),
				user:    &User{Name: "bob", Password: pw},
				wantErr: true,
				want:    nil,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServiceRead(t *testing.T) {
	type test struct {
		testcase string
		service  Service
		id       ID
		wantErr  bool
		want     *User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got, err := tt.service.Read(tt.id)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 1",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				id:      1,
				wantErr: false,
				want:    &User{ID: 1, Name: "bob", Password: pw},
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error 1",
				service: NewService(&mockRepository{
					user: &User{ID: 2, Name: "alice", Password: pw},
					err:  fmt.Errorf("test error"),
				}),
				id:      2,
				wantErr: true,
				want:    nil,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServiceUpdate(t *testing.T) {
	type test struct {
		testcase string
		service  Service
		user     *User
		wantErr  bool
		want     *User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got, err := tt.service.Update(tt.user)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 1",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				user:    &User{Name: "bob", Password: pw},
				wantErr: false,
				want:    &User{ID: 1, Name: "bob", Password: pw},
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error invalid user",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				user:    &User{Name: "", Password: pw},
				wantErr: true,
				want:    nil,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error repository",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  fmt.Errorf("test error"),
				}),
				user:    &User{Name: "bob", Password: pw},
				wantErr: true,
				want:    nil,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServiceDelete(t *testing.T) {
	type test struct {
		testcase string
		service  Service
		id       ID
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		err := tt.service.Delete(tt.id)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Errorf("want-err=%v, err=%v.", tt.wantErr, err)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 1",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  nil,
				}),
				id:      1,
				wantErr: false,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "error 1",
				service: NewService(&mockRepository{
					user: &User{ID: 1, Name: "bob", Password: pw},
					err:  fmt.Errorf("test error"),
				}),
				id:      1,
				wantErr: true,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}
