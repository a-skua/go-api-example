package request

import (
	"api.example.com/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestUserCreate(t *testing.T) {
	tests := []struct {
		in       *http.Request
		password string
		want     *entity.User
	}{
		{
			httptest.NewRequest(
				"POST",
				"http://api.example.com/user",
				strings.NewReader(`{"user":{"name":"Bob","password":"qwerty"}}`),
			),
			"qwerty",
			&entity.User{Name: "Bob", Password: nil}, // TODO Password は DeepEqual の対象としない
		},
	}

	for _, tt := range tests {
		got, err := UserCreate(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		ok := got.Password.Verify(tt.password)
		if !ok {
			t.Fatalf("request.UserCreate: invalid password")
		}
		got.Password = nil // TODO テスト方法を考える

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("request.UserCreate: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestUserRead(t *testing.T) {
	tests := []struct {
		in   *http.Request
		want entity.UserID
	}{
		{
			mux.SetURLVars(httptest.NewRequest(
				"GET",
				"http://api.example.com/user/1",
				nil,
			), map[string]string{"user_id": "1"}),
			1,
		},
	}

	for _, tt := range tests {
		got, err := UserRead(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if tt.want != got {
			t.Fatalf("request.UserRead: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestFailedUserRead(t *testing.T) {
	tests := []struct {
		in *http.Request
	}{
		{
			httptest.NewRequest(
				"GET",
				"http://api.example.com/user/1",
				nil,
			),
		},
	}

	for _, tt := range tests {
		_, err := UserRead(tt.in)
		if err == nil {
			t.Fatalf("Expect Error")
		}
		t.Log(err)
	}
}
