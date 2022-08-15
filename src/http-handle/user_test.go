package handle

import (
	"api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// mock
type userServer struct {
	user *user.User
	err  error
	// flags
	create, read, update, delete bool
}

func (s *userServer) Create(*user.User) (*user.User, error) {
	if s.create {
		return s.user, s.err
	}

	panic("invalid Create")
}

func (s *userServer) Read(user.ID) (*user.User, error) {
	if s.read {
		return s.user, s.err
	}

	panic("invalid Read")
}

func (s *userServer) Update(*user.User) (*user.User, error) {
	if s.update {
		return s.user, s.err
	}

	panic("invalid Update")
}

func (s *userServer) Delete(user.ID) error {
	if s.delete {
		return s.err
	}

	panic("invalid Delete")
}

// test
func TestUserHandler_create(t *testing.T) {
	type args struct {
		url  string
		body []byte
	}

	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		args
		server user.Server
		want   want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer(tt.args.body))
			w := httptest.NewRecorder()

			s := newServices()
			s.User = tt.server
			New(s).ServeHTTP(w, r)

			got := w.Result()
			defer got.Body.Close()

			gotBody, _ := io.ReadAll(got.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := got.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := got.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				url:  "/user",
				body: []byte(`{"user":{"name":"bob","password":"qwerty"}}`),
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				create: true,
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"bob","password":"*****"}}` + "\n"),
			},
		},
		{
			testcase: "invalid request-body",
			args: args{
				url:  "/user",
				body: []byte(`id=1&name=bob&password=qwerty`),
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				create: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
		{
			testcase: "failed server-create",
			args: args{
				url:  "/user",
				body: []byte(`{"user":{"name":"bob","password":"qwerty"}}`),
			},
			server: &userServer{
				err:    errors.New("internal server error"),
				create: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserHandler_read(t *testing.T) {
	type args struct {
		url  string
		body []byte
	}

	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		args
		server user.Server
		want   want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.url, bytes.NewBuffer(tt.args.body))
			w := httptest.NewRecorder()

			s := newServices()
			s.User = tt.server
			New(s).ServeHTTP(w, r)

			got := w.Result()
			defer got.Body.Close()

			gotBody, _ := io.ReadAll(got.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := got.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := got.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				url:  "/user/1",
				body: []byte{},
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				read: true,
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"bob","password":"*****"}}` + "\n"),
			},
		},
		{
			testcase: "invalid user_id",
			args: args{
				url:  "/user/xxx",
				body: []byte{},
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				read: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
		{
			testcase: "failed server-read",
			args: args{
				url:  "/user/1",
				body: []byte{},
			},
			server: &userServer{
				err:  errors.New("internal server error"),
				read: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserHandler_update(t *testing.T) {
	type args struct {
		url  string
		body []byte
	}

	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		args
		server user.Server
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPut, tt.url, bytes.NewBuffer(tt.args.body))
			w := httptest.NewRecorder()

			s := newServices()
			s.User = tt.server
			New(s).ServeHTTP(w, r)

			got := w.Result()
			defer got.Body.Close()

			gotBody, _ := io.ReadAll(got.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := got.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := got.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				url:  "/user/1",
				body: []byte(`{"user":{"name":"bob","password":"*****"}}`),
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				update: true,
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{"id":1,"name":"bob","password":"*****"}}` + "\n"),
			},
		},
		{
			testcase: "invalid user_id",
			args: args{
				url:  "/user/xxx",
				body: []byte{},
			},
			server: &userServer{
				user: &user.User{
					ID:       1,
					Name:     "bob",
					Password: password.FromHash([]byte("qwerty")),
				},
				update: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
		{
			testcase: "failed server-read",
			args: args{
				url:  "/user/1",
				body: []byte(`{"user":{"name":"bob","password":"*****"}}`),
			},
			server: &userServer{
				err:    errors.New("internal server error"),
				update: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestUserHandler_delete(t *testing.T) {
	type args struct {
		url  string
		body []byte
	}

	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		args
		server user.Server
		want   want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodDelete, tt.url, bytes.NewBuffer(tt.args.body))
			w := httptest.NewRecorder()

			s := newServices()
			s.User = tt.server
			New(s).ServeHTTP(w, r)

			got := w.Result()
			defer got.Body.Close()

			gotBody, _ := io.ReadAll(got.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%v, got=%v.", tt.want.body, gotBody)
			}

			gotContentType := got.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := got.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				url:  "/user/1",
				body: []byte{},
			},
			server: &userServer{
				err:    nil,
				delete: true,
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"user":{}}` + "\n"),
			},
		},
		{
			testcase: "invalid user_id",
			args: args{
				url:  "/user/xxx",
				body: []byte{},
			},
			server: &userServer{
				err:    nil,
				delete: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
		{
			testcase: "failed server-read",
			args: args{
				url:  "/user/1",
				body: []byte{},
			},
			server: &userServer{
				err:    errors.New("internal server error"),
				delete: true,
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
