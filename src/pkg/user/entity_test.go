package user

import (
	"reflect"
	"testing"
)

func TestNewPassword(t *testing.T) {
	type test struct {
		testcase string
		in       []byte
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got, err := NewPassword(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		pw, ok := got.(*password)
		if !ok {
			t.Fatalf("invalid type: want=%T, got=%T.", pw, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in:       []byte("password"),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNewPasswordFromHash(t *testing.T) {
	type test struct {
		testcase string
		in       []byte
		want     Password
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := NewPasswordFromHash(tt.in)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in:       []byte("qwerty"),
			want: &password{
				hash:   []byte("qwerty"),
				length: 8,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPasswordString(t *testing.T) {
	type test struct {
		testcase string
		password Password
		want     string
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.password.String()
		if tt.want != got {
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
				password: pw,
				want:     "*****",
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("qwerty"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 2",
				password: pw,
				want:     "*****",
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPasswordVerify(t *testing.T) {
	type test struct {
		testcase string
		password Password
		in       []byte
		want     bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.password.Verify(tt.in)
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			password, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 1",
				password: password,
				in:       []byte("password"),
				want:     true,
			}
		}(),
		func() test {
			password, err := NewPassword([]byte("qwerty"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success 2",
				password: password,
				in:       []byte("qwerty"),
				want:     true,
			}
		}(),
		func() test {
			password, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "failed 1",
				password: password,
				in:       []byte("qwerty"),
				want:     false,
			}
		}(),
		func() test {
			password, err := NewPassword([]byte("qwerty"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "failed 2",
				password: password,
				in:       []byte("password"),
				want:     false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPasswordLength(t *testing.T) {
	type test struct {
		testcase string
		password Password
		want     int
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.password.Length()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "length 1",
			password: &password{
				length: 1,
			},
			want: 1,
		},
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "length 8",
				password: pw,
				want:     8,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("qwerty"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "length 6",
				password: pw,
				want:     6,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPasswordHash(t *testing.T) {
	type test struct {
		testcase string
		password Password
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.password.Hash()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "case 1",
			password: &password{hash: []byte("dummy hash!")},
			want:     []byte("dummy hash!"),
		},
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			hash := func() []byte {
				pw, ok := pw.(*password)
				if !ok {
					t.Fatalf("invalid type: want=%T.", pw)
				}
				return pw.hash
			}()

			return test{
				testcase: "case 2",
				password: pw,
				want:     hash,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNew(t *testing.T) {
	type test struct {
		testcase string
		name     string
		password Password
		want     *User
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := New(tt.name, tt.password)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("qwerty"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "case 1",
				name:     "bob",
				password: pw,
				want: &User{
					ID:       0,
					Name:     "bob",
					Password: pw,
				},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserValid(t *testing.T) {
	type test struct {
		testcase string
		user     *User
		want     bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.user.Valid()
		if tt.want != got {
			t.Fatalf("want=%v, got%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "success",
				user: &User{
					Name:     "bob",
					Password: pw,
				},
				want: true,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte(""))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "empty password",
				user: &User{
					Name:     "bob",
					Password: pw,
				},
				want: false,
			}
		}(),
		func() test {
			pw, err := NewPassword([]byte("password"))
			if err != nil {
				panic(err)
			}

			return test{
				testcase: "empty name",
				user: &User{
					Name:     "",
					Password: pw,
				},
				want: false,
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}
