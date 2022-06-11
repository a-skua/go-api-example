package response

import (
	"api.example.com/pkg/user"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// mock Password
type mockPassword string

func (pw mockPassword) Verify(plain string) bool {
	return pw == mockPassword(plain)
}

func (pw mockPassword) Length() int {
	return len(pw)
}

func (mockPassword) String() string {
	return user.PasswordString
}

func (pw mockPassword) Hash() []byte {
	return []byte(pw)
}

// test
func TestUserCreate(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		user     *user.User
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := UserCreate(w, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := res.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := res.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			user: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: mockPassword("password"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
			},
		},
		{
			user: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: mockPassword("qwerty"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserRead(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		user     *user.User
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := UserRead(w, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := res.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := res.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			user: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: mockPassword("password"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
			},
		},
		{
			user: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: mockPassword("qwerty"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserUpdate(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		user     *user.User
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := UserUpdate(w, tt.user)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := res.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := res.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			user: &user.User{
				ID:       1,
				Name:     "Bob",
				Password: mockPassword("password"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"Bob","password":"*****"}}` + "\n"),
			},
		},
		{
			user: &user.User{
				ID:       2,
				Name:     "Alice",
				Password: mockPassword("password"),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":2,"name":"Alice","password":"*****"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserDelete(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := UserDelete(w)
			if hasErr := err != nil; tt.wantErr != hasErr {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := res.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := res.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
