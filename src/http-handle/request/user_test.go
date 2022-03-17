package request

import (
	"api.example.com/pkg/user"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestUserCreate(t *testing.T) {
	type test struct {
		testcase string
		in       *http.Request
		password []byte
		want     *user.User
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		got, err := UserCreate(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
		}

		if !tt.wantErr {
			ok := got.Password.Verify(tt.password)
			if !ok {
				t.Fatalf("invalid password=%v.", tt.password)
			}
			// NOTE
			// password 検証ができたので比較のために初期化
			got.Password = nil
			tt.want.Password = nil
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in: httptest.NewRequest(
				"POST",
				"http://api.example.com/user",
				strings.NewReader(`{"user":{"name":"Bob","password":"qwerty"}}`),
			),
			password: []byte("qwerty"),
			want:     &user.User{Name: "Bob", Password: nil},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		testcase string
		in       *http.Request
		want     user.ID
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		got, err := UserRead(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
		}

		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in: mux.SetURLVars(httptest.NewRequest(
				"GET",
				"http://api.example.com/user/1",
				nil,
			), map[string]string{"user_id": "1"}),
			want:    1,
			wantErr: false,
		},
		{
			testcase: "error",
			in: httptest.NewRequest(
				"GET",
				"http://api.example.com/user/1",
				nil,
			),
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type test struct {
		testcase string
		in       *http.Request
		password []byte
		want     *user.User
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		got, err := UserUpdate(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
		}

		if !tt.wantErr {
			ok := got.Password.Verify(tt.password)
			if !ok {
				t.Fatalf("invalid password=%v.", tt.password)
			}
			// NOTE
			// password 検証ができたので比較のために初期化
			got.Password = nil
			tt.want.Password = nil
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in: mux.SetURLVars(httptest.NewRequest(
				"GET",
				"http://api.example.com/user/1",
				strings.NewReader(`{"user":{"name":"Bob","password":"qwerty"}}`),
			), map[string]string{"user_id": "1"}),
			password: []byte("qwerty"),
			want: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		testcase string
		in       *http.Request
		want     user.ID
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		got, err := UserDelete(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in: mux.SetURLVars(httptest.NewRequest(
				"GET",
				"http://api.example.com/user/2",
				nil,
			), map[string]string{"user_id": "2"}),
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
