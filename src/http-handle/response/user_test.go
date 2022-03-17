package response

import (
	"api.example.com/pkg/user"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

// mock Password
type plainPassword string

func (pw plainPassword) Verify(plain []byte) bool {
	return pw == plainPassword(plain)
}

func (pw plainPassword) Length() int {
	return len(pw)
}

func (plainPassword) String() string {
	return "*****"
}

func (pw plainPassword) Hash() []byte {
	return []byte(pw)
}

func TestUserCreate(t *testing.T) {
	type test struct {
		testcase string
		in       *user.User
		wantErr  bool
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := UserCreate(w, tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "saccess bob",
			in: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
		},
		{
			testcase: "saccess alice",
			in: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: plainPassword("qwerty"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		testcase string
		in       *user.User
		wantErr  bool
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := UserRead(w, tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success bob",
			in: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
		},
		{
			testcase: "success alice",
			in: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: plainPassword("qwerty"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type test struct {
		testcase string
		in       *user.User
		wantErr  bool
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := UserUpdate(w, tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success bob",
			in: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: plainPassword("password"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
		},
		{
			testcase: "success alice",
			in: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: plainPassword("password"),
			},
			wantErr: false,
			want:    []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		testcase string
		wantErr  bool
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := UserDelete(w)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			wantErr:  false,
			want:     []byte(`{"user":{}}` + "\n"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
