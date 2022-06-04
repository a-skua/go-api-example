package user

import (
	"fmt"
	"reflect"
	"testing"
)

// mock
type mockRepository struct {
	create, read, update, delete bool
	user                         *User
	err                          error
}

func (r *mockRepository) UserCreate(*User) (*User, error) {
	if r.create {
		return r.user, r.err
	} else {
		return nil, fmt.Errorf("failed create")
	}
}

func (r *mockRepository) UserRead(ID) (*User, error) {
	if r.read {
		return r.user, r.err
	} else {
		return nil, fmt.Errorf("failed read")
	}
}

func (r *mockRepository) UserUpdate(*User) (*User, error) {
	if r.update {
		return r.user, r.err
	} else {
		return nil, fmt.Errorf("failed update")
	}
}

func (r *mockRepository) UserDelete(ID) error {
	if r.delete {
		return r.err
	} else {
		return fmt.Errorf("failed delete")
	}
}

// test
func TestNewServer(t *testing.T) {
	type test struct {
		testcase   string
		repository Repository
		want       Server
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := NewServer(tt.repository)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			repository: &mockRepository{},
			want: &server{
				repository: &mockRepository{},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Create(t *testing.T) {
	type test struct {
		testcase string
		server   Server
		user     *User
		wantErr  bool
		want     *User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Create(tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			t.Log(err)

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			server: NewServer(&mockRepository{
				create: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"password"},
				},
				err: nil,
			}),
			user:    New("bob", &mockPassword{"password"}),
			wantErr: false,
			want: &User{
				ID:       1,
				Name:     "bob",
				Password: &mockPassword{"password"},
			},
		},
		{
			testcase: "invalid user",
			server: NewServer(&mockRepository{
				create: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"password"},
				},
				err: nil,
			}),
			user:    New("bob", &mockPassword{"qwerty"}),
			wantErr: true,
			want:    nil,
		},
		{
			testcase: "repository error",
			server: NewServer(&mockRepository{
				create: true,
				user:   nil,
				err:    fmt.Errorf("internal server error"),
			}),
			user:    New("bob", &mockPassword{"password"}),
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Read(t *testing.T) {
	type test struct {
		testcase string
		server   Server
		id       ID
		wantErr  bool
		want     *User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Read(tt.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			t.Log(err)

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			server: NewServer(&mockRepository{
				read: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"qwerty"},
				},
				err: nil,
			}),
			id:      1,
			wantErr: false,
			want: &User{
				ID:       1,
				Name:     "bob",
				Password: &mockPassword{"qwerty"},
			},
		},
		{
			testcase: "invalid user_id",
			server: NewServer(&mockRepository{
				read: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"qwerty"},
				},
				err: nil,
			}),
			id:      0,
			wantErr: true,
			want:    nil,
		},
		{
			testcase: "repository error",
			server: NewServer(&mockRepository{
				read: true,
				user: nil,
				err:  fmt.Errorf("internal server error"),
			}),
			id:      1,
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Update(t *testing.T) {
	type test struct {
		testcase string
		server   Server
		user     *User
		wantErr  bool
		want     *User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Update(tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			t.Log(err)

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			server: NewServer(&mockRepository{
				update: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"password"},
				},
				err: nil,
			}),
			user: &User{
				ID:       1,
				Name:     "bob",
				Password: &mockPassword{"password"},
			},
			wantErr: false,
			want: &User{
				ID:       1,
				Name:     "bob",
				Password: &mockPassword{"password"},
			},
		},
		{
			testcase: "invalid user",
			server: NewServer(&mockRepository{
				update: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"password"},
				},
				err: nil,
			}),
			user: &User{
				ID:       1,
				Name:     "",
				Password: &mockPassword{"password"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			testcase: "invalid user_id",
			server: NewServer(&mockRepository{
				update: true,
				user: &User{
					ID:       1,
					Name:     "bob",
					Password: &mockPassword{"password"},
				},
				err: nil,
			}),
			user: &User{
				ID:       0,
				Name:     "bob",
				Password: &mockPassword{"password"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			testcase: "repository error",
			server: NewServer(&mockRepository{
				update: true,
				user:   nil,
				err:    fmt.Errorf("internal server error"),
			}),
			user: &User{
				ID:       1,
				Name:     "bob",
				Password: &mockPassword{"password"},
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Delete(t *testing.T) {
	type test struct {
		testcase string
		server   Server
		id       ID
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			err := tt.server.Delete(tt.id)
			if tt.wantErr != (err != nil) {
				t.Errorf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			t.Log(err)
		})
	}

	tests := []*test{
		{
			server: NewServer(&mockRepository{
				delete: true,
				err:    nil,
			}),
			id:      1,
			wantErr: false,
		},
		{
			testcase: "invalid user_id",
			server: NewServer(&mockRepository{
				delete: true,
				err:    nil,
			}),
			id:      0,
			wantErr: true,
		},
		{
			testcase: "repository error",
			server: NewServer(&mockRepository{
				delete: true,
				err:    fmt.Errorf("internal server error"),
			}),
			id:      1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
