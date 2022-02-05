package handle

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func testHeader(res *http.Response) error {
	want := "application/json"
	got := res.Header.Get("Content-Type")
	if want != got {
		return fmt.Errorf("Header Content-Type: want=%v, got=%v.", want, got)
	}
	return nil
}

func TestUserCreate(t *testing.T) {
	pw, _ := entity.NewPassword("password")

	tests := []struct {
		handle http.Handler
		in     *http.Request
		want   []byte
	}{
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: pw,
			}}),
			in: httptest.NewRequest(
				"POST",
				"http://api.example.com/user",
				strings.NewReader(``),
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{Error: errors.New("Repository Error")}),
			in: httptest.NewRequest(
				"POST",
				"http://api.example.com/user",
				strings.NewReader(`{"user":{"name":"Bob","password":"password"}}`),
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       1,
				Name:     "Bob",
				Password: pw,
			}}),
			in: httptest.NewRequest(
				"POST",
				"http://api.example.com/user",
				strings.NewReader(`{"user":{"name":"Bob","password":"password"}}`),
			),
			want: []byte(`{"user":{"id":1,"name":"Bob","password":"*****","companies":[]}}` + "\n"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		tt.handle.ServeHTTP(w, tt.in)

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		err := testHeader(res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%s, got=%s.", tt.want, got)
		}
	}
}

func TestUserRead(t *testing.T) {
	pw, _ := entity.NewPassword("querty")

	tests := []struct {
		handle  http.Handler
		in      *http.Request
		xUserID string
		want    []byte
	}{
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Alice",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "",
			in: httptest.NewRequest(
				"GET",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle:  New(&repository.Mock{Error: errors.New("Repository Error")}),
			xUserID: "2",
			in: httptest.NewRequest(
				"GET",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Alice",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "2",
			in: httptest.NewRequest(
				"GET",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"user":{"id":2,"name":"Alice","password":"*****","companies":[{"id":1,"name":"GREATE COMPANY"}]}}` + "\n"),
		},
	}

	for _, tt := range tests {
		tt.in.Header.Set("X-User-Id", tt.xUserID)

		w := httptest.NewRecorder()
		tt.handle.ServeHTTP(w, tt.in)

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		err := testHeader(res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%s, got=%s.", tt.want, got)
		}
	}
}

func TestUserUpdate(t *testing.T) {
	pw, _ := entity.NewPassword("12345678")

	tests := []struct {
		handle  http.Handler
		in      *http.Request
		xUserID string
		want    []byte
	}{
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Bob",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "2",
			in: httptest.NewRequest(
				"PUT",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Bob",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "",
			in: httptest.NewRequest(
				"PUT",
				"http://api.example.com/user/2",
				strings.NewReader(`{"user":{"name":"Bob","password":"qwerty"}}`),
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Bob",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "2",
			in: httptest.NewRequest(
				"PUT",
				"http://api.example.com/user/2",
				strings.NewReader(`{"user":{"name":"Bob","password":"qwerty"}}`),
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle: New(&repository.Mock{User: &entity.User{
				ID:       2,
				Name:     "Bob",
				Password: pw,
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			}}),
			xUserID: "2",
			in: httptest.NewRequest(
				"PUT",
				"http://api.example.com/user/2",
				strings.NewReader(`{"user":{"name":"Bob","password":"12345678"}}`),
			),
			want: []byte(`{"user":{"id":2,"name":"Bob","password":"*****","companies":[{"id":1,"name":"GREATE COMPANY"}]}}` + "\n"),
		},
	}

	for _, tt := range tests {
		tt.in.Header.Set("X-User-Id", tt.xUserID)

		w := httptest.NewRecorder()
		tt.handle.ServeHTTP(w, tt.in)

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		err := testHeader(res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%s, got=%s.", tt.want, got)
		}
	}
}

func TestUserDelete(t *testing.T) {
	tests := []struct {
		handle  http.Handler
		in      *http.Request
		xUserID string
		want    []byte
	}{
		{
			handle:  New(&repository.Mock{Error: nil}),
			xUserID: "",
			in: httptest.NewRequest(
				"DELETE",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle:  New(&repository.Mock{Error: errors.New("Repository Error")}),
			xUserID: "2",
			in: httptest.NewRequest(
				"DELETE",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
		{
			handle:  New(&repository.Mock{Error: nil}),
			xUserID: "2",
			in: httptest.NewRequest(
				"DELETE",
				"http://api.example.com/user/2",
				nil,
			),
			want: []byte(`{"user":{}}` + "\n"),
		},
	}

	for _, tt := range tests {
		tt.in.Header.Set("X-User-Id", tt.xUserID)

		w := httptest.NewRecorder()
		tt.handle.ServeHTTP(w, tt.in)

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		err := testHeader(res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%s, got=%s.", tt.want, got)
		}
	}
}
