package password

import (
	"api.example.com/pkg/user"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type test struct {
		testcase string
		in       string
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := New(tt.in)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			if got == nil {
				t.Fatalf("NewPassword return nil")
			}
		})
	}

	tests := []*test{
		{
			in:      "password",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestFromHash(t *testing.T) {
	type test struct {
		testcase string
		in       []byte
		want     user.Password
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := FromHash(tt.in)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			in: []byte("qwerty"),
			want: &password{
				hash:   []byte("qwerty"),
				length: 0,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPassword_String(t *testing.T) {
	type test struct {
		testcase string
		password user.Password
		want     string
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.password.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			password: &password{
				hash: []byte("password"),
			},
			want: user.PasswordString,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPassword_Verify(t *testing.T) {
	type test struct {
		testcase string
		password user.Password
		in       string
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.password.Verify(tt.in)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		func() *test {
			password, err := New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				password: password,
				in:       "password",
				want:     true,
			}
		}(),
		func() *test {
			password, err := New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				password: password,
				in:       "qwerty",
				want:     false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPassword_Length(t *testing.T) {
	type test struct {
		testcase string
		password user.Password
		want     int
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.password.Length()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		func() *test {
			pw, err := New("password")
			if err != nil {
				panic(err)
			}

			return &test{
				password: pw,
				want:     8,
			}
		}(),
		{
			testcase: "length 1",
			password: &password{
				length: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPassword_Hash(t *testing.T) {
	type test struct {
		testcase string
		password user.Password
		want     []byte
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.password.Hash()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			password: &password{
				hash: []byte("password_hash!"),
			},
			want: []byte("password_hash!"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
