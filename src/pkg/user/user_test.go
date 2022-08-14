package user

import (
	"reflect"
	"strings"
	"testing"
)

// mock password
type password struct {
	hash string
}

func newPassword(plain string) Password {
	return &password{plain}
}

func (pw *password) String() string {
	return "*****"
}

func (pw *password) Verify(plain string) bool {
	return pw.hash == plain
}

func (pw *password) Length() int {
	return len(pw.hash)
}

func (pw *password) Hash() []byte {
	return []byte(pw.hash)
}

func TestID_valid(t *testing.T) {
	type test struct {
		name string
		id   ID
		want bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.id.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "id ≥ 1",
			id:   1,
			want: true,
		},
		{
			name: "id ≤ 1",
			id:   0,
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestName_valid(t *testing.T) {
	type test struct {
		testname string
		username Name
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testname, func(t *testing.T) {
			got := tt.username.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			testname: "length 1",
			username: "1",
			want:     true,
		},
		{
			testname: "length 0",
			username: "",
			want:     false,
		},
		{
			testname: "length 255",
			username: Name(strings.Repeat("1", 255)),
			want:     true,
		},
		{
			testname: "length 256",
			username: Name(strings.Repeat("1", 256)),
			want:     false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestValidPassword(t *testing.T) {
	type test struct {
		name     string
		password Password
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := validPassword(tt.password)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:     "length 8",
			password: newPassword("password"),
			want:     true,
		},
		{
			name:     "length 7",
			password: newPassword("1234567"),
			want:     false,
		},
		{
			name:     "length 255",
			password: newPassword(strings.Repeat("1", 255)),
			want:     true,
		},
		{
			name:     "length 256",
			password: newPassword(strings.Repeat("1", 256)),
			want:     false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		name     Name
		password Password
	}

	type test struct {
		name string
		args args
		want *User
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.name, tt.args.password)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				name:     "foo",
				password: newPassword("bar"),
			},
			want: &User{
				Name:     "foo",
				Password: newPassword("bar"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_validCreate(t *testing.T) {
	type test struct {
		name string
		user *User
		want bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.user.validCreate()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "valid",
			user: &User{
				Name:     "Bob",
				Password: newPassword("password"),
			},
			want: true,
		},
		{
			name: "invalid user.name",
			user: &User{
				Name:     "",
				Password: newPassword("password"),
			},
			want: false,
		},
		{
			name: "invalid user.password",
			user: &User{
				Name:     "Bob",
				Password: newPassword("1234567"),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_validUpdate(t *testing.T) {
	type test struct {
		name string
		user *User
		want bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.user.validUpdate()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "valid",
			user: &User{
				ID:       1,
				Name:     "Bob",
				Password: newPassword("password"),
			},
			want: true,
		},
		{
			name: "invalid user.id",
			user: &User{
				ID:       0,
				Name:     "Bob",
				Password: newPassword("password"),
			},
			want: false,
		},
		{
			name: "invalid user.name",
			user: &User{
				ID:       1,
				Name:     "",
				Password: newPassword("password"),
			},
			want: false,
		},
		{
			name: "invalid user.password",
			user: &User{
				ID:       1,
				Name:     "Bob",
				Password: newPassword(""),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
