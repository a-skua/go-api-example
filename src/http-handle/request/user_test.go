package request

import (
	"api.example.com/pkg/user"
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUserCreate(t *testing.T) {
	type test struct {
		testcase string
		body     []byte
		password string
		want     *user.User
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest("POST", "http://api.example.com/user", bytes.NewBuffer(tt.body))

			got, err := UserCreate(r)
			if tt.wantErr != (err != nil) {
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
		})
	}

	tests := []*test{
		{
			body:     []byte(`{"user":{"name":"Bob","password":"qwerty"}}`),
			password: "qwerty",
			want:     &user.User{Name: "Bob", Password: nil},
			wantErr:  false,
		},
		{
			body:    []byte(""),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type test struct {
		testcase string
		url      string
		want     user.ID
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tt.url, nil)

			var (
				got user.ID
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = UserRead(r)
			})
			router.ServeHTTP(w, r)

			if hasErr := err != nil; tt.wantErr != hasErr {
				t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			url:     "http://api.example.com/user/1",
			want:    1,
			wantErr: false,
		},
		{
			url:     "http://api.example.com/user/foo",
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
		url      string
		body     []byte
		password string
		want     *user.User
		wantErr  bool
	}

	do := func(tt test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", tt.url, bytes.NewBuffer(tt.body))

			var (
				got *user.User
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = UserUpdate(r)
			})
			router.ServeHTTP(w, r)

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
		})
	}

	tests := []test{
		{
			url:      "http://api.example.com/user/1",
			body:     []byte(`{"user":{"name":"Bob","password":"qwerty"}}`),
			password: "qwerty",
			want: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: nil,
			},
			wantErr: false,
		},
		{
			testcase: "empty body",
			url:      "http://api.example.com/user/1",
			body:     []byte(""),
			want:     nil,
			wantErr:  true,
		},
		{
			testcase: "invalid user_id",
			url:      "http://api.example.com/user/xxx",
			body:     []byte(`{"user":{"name":"Bob","password":"qwerty"}}`),
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type test struct {
		testcase string
		url      string
		want     user.ID
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tt.url, nil)

			var (
				got user.ID
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = UserDelete(r)
			})
			router.ServeHTTP(w, r)

			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []test{
		{
			url:     "http://api.example.com/user/2",
			want:    2,
			wantErr: false,
		},
		{
			url:     "http://api.example.com/user/xxx",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
