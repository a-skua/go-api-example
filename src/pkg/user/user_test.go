package user

import (
	"reflect"
	"testing"
)

// mock
type mockPassword struct {
	hash string
}

func (pw *mockPassword) String() string {
	return PasswordString
}

func (pw *mockPassword) Verify(plain string) bool {
	return pw.hash == plain
}

func (pw *mockPassword) Length() int {
	return len(pw.hash)
}

func (pw *mockPassword) Hash() []byte {
	return []byte(pw.hash)
}

// test
func TestID_Valid(t *testing.T) {
	type test struct {
		testcase string
		id       ID
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.id.Valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			id:   0,
			want: false,
		},
		{
			id:   1,
			want: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestName_valid(t *testing.T) {
	type test struct {
		testcase string
		name     Name
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.name.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "1",
			want: true,
		},
		{
			name: "",
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestValidPassword(t *testing.T) {
	type test struct {
		testcase string
		pw       Password
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := validPassword(tt.pw)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			pw:   &mockPassword{"password"},
			want: true,
		},
		{
			pw:   &mockPassword{"1234567"},
			want: false,
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
		testcase string
		args
		want *User
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := New(tt.name, tt.password)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				name:     "foo",
				password: &mockPassword{},
			},
			want: &User{
				Name:     "foo",
				Password: &mockPassword{},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUser_valid(t *testing.T) {
	type test struct {
		testcase string
		user     *User
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.user.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			user: &User{
				Name:     "foo",
				Password: &mockPassword{"password"},
			},
			want: true,
		},
		{
			testcase: "invalid name",
			user: &User{
				Name:     "",
				Password: &mockPassword{"password"},
			},
			want: false,
		},
		{
			testcase: "invalid password",
			user: &User{
				Name:     "foo",
				Password: &mockPassword{""},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
