package user

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// mock
type repository struct {
	user *User
	err  error
	// flags
	create, read, update, delete bool
}

func (r *repository) UserCreate(*User) (*User, error) {
	if r.create {
		return r.user, r.err
	}
	return nil, fmt.Errorf("failed create")
}

func (r *repository) UserRead(ID) (*User, error) {
	if r.read {
		return r.user, r.err
	}
	return nil, fmt.Errorf("failed read")
}

func (r *repository) UserUpdate(*User) (*User, error) {
	if r.update {
		return r.user, r.err
	}
	return nil, fmt.Errorf("failed update")
}

func (r *repository) UserDelete(ID) error {
	if r.delete {
		return r.err
	}
	return fmt.Errorf("failed delete")
}

// test
func TestNewServer(t *testing.T) {
	type args struct {
		repository Repository
	}

	type test struct {
		name string
		args args
		want Server
	}

	do := func(tt *test) {
		got := NewServer(tt.args.repository)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []*test{
		{
			name: "true",
			args: args{
				repository: &repository{},
			},
			want: &server{
				repository: &repository{},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Create(t *testing.T) {
	type args struct {
		user *User
	}

	type test struct {
		name    string
		server  Server
		args    args
		want    *User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.Create(tt.args.user)
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
			name: "true",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			}),
			args: args{
				user: New("Bob", newPassword("password")),
			},
			want: &User{
				ID:        1,
				Name:      "Bob",
				Password:  newPassword("password"),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "invalid user.name",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			}),
			args: args{
				user: New("", newPassword("password")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid user.name",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			}),
			args: args{
				user: New("", newPassword("password")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid user.password",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				create: true,
			}),
			args: args{
				user: New("Bob", newPassword("")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed create",
			server: NewServer(&repository{
				err:    errors.New("inernal server error"),
				create: true,
			}),
			args: args{
				user: New("Bob", newPassword("password")),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Read(t *testing.T) {
	type args struct {
		id ID
	}

	type test struct {
		name    string
		server  Server
		args    args
		want    *User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.Read(tt.args.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want=%v, got=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				read: true,
			}),
			args: args{
				id: 1,
			},
			want: &User{
				ID:        1,
				Name:      "Bob",
				Password:  newPassword("password"),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "invalid user.id",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				read: true,
			}),
			args: args{
				id: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed read",
			server: NewServer(&repository{
				err:  errors.New("internal server error"),
				read: true,
			}),
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Update(t *testing.T) {
	type args struct {
		user *User
	}

	type test struct {
		name    string
		server  Server
		args    args
		want    *User
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.Update(tt.args.user)
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
			name: "true",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			}),
			args: args{
				user: &User{
					ID:       1,
					Name:     "Bob",
					Password: newPassword("password"),
				},
			},
			want: &User{
				ID:        1,
				Name:      "Bob",
				Password:  newPassword("password"),
				UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "invalid user.id",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			}),
			args: args{
				user: &User{
					ID:       0,
					Name:     "Bob",
					Password: newPassword("password"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid user.name",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			}),
			args: args{
				user: &User{
					ID:       1,
					Name:     "",
					Password: newPassword("password"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid user.password",
			server: NewServer(&repository{
				user: &User{
					ID:        1,
					Name:      "Bob",
					Password:  newPassword("password"),
					UpdatedAt: time.Date(2022, 8, 9, 12, 34, 56, 0, time.UTC),
				},
				update: true,
			}),
			args: args{
				user: &User{
					ID:       1,
					Name:     "Bob",
					Password: newPassword(""),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed update",
			server: NewServer(&repository{
				err:    errors.New("internal server error"),
				update: true,
			}),
			args: args{
				user: &User{
					ID:       1,
					Name:     "Bob",
					Password: newPassword("password"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Delete(t *testing.T) {
	type args struct {
		id ID
	}

	type test struct {
		name    string
		server  Server
		args    args
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.server.Delete(tt.args.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			server: NewServer(&repository{
				err:    nil,
				delete: true,
			}),
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid user.id",
			server: NewServer(&repository{
				err:    nil,
				delete: true,
			}),
			args: args{
				id: 0,
			},
			wantErr: true,
		},
		{
			name: "failed delete",
			server: NewServer(&repository{
				err:    errors.New("internal server error"),
				delete: true,
			}),
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
