package handle

import (
	"api.example.com/pkg/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthUser(t *testing.T) {
	tests := []struct {
		in      *http.Request
		xUserID string
		want    entity.UserID
	}{
		{
			httptest.NewRequest("GET", "http://api.example.com/user/1", nil),
			"1",
			1,
		},
	}

	for _, tt := range tests {
		tt.in.Header.Set("X-User-Id", tt.xUserID)
		got, err := authUser(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if tt.want != got {
			t.Fatalf("handle.authUser: want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestFailedAuthUser(t *testing.T) {
	tests := []struct {
		in *http.Request
	}{
		{
			httptest.NewRequest("GET", "http://api.example.com/user/1", nil),
		},
	}

	for _, tt := range tests {
		_, err := authUser(tt.in)
		if err == nil {
			t.Fatalf("Expect Errro")
		}
		t.Log(err)
	}
}
