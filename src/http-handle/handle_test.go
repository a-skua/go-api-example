package handle

import (
	"api.example.com/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewAuth(t *testing.T) {
	type test struct {
		testcase string
		in       *http.Request
		xUserID  string
		wantErr  bool
		want     auth
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		tt.in.Header.Set("X-User-Id", tt.xUserID)
		got, err := newAuth(tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatal(err)
		}

		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in:       httptest.NewRequest("GET", "http://api.example.com/user/1", nil),
			xUserID:  "1",
			wantErr:  false,
			want:     1,
		},
		{
			testcase: "error",
			in:       httptest.NewRequest("GET", "http://api.example.com/user/1", nil),
			xUserID:  "",
			wantErr:  true,
			want:     0,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestAuthVerify(t *testing.T) {
	type test struct {
		testcase string
		auth     auth
		in       user.ID
		want     bool
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		got := tt.auth.verify(tt.in)
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			auth:     1,
			in:       1,
			want:     true,
		},
		{
			testcase: "failed",
			auth:     1,
			in:       2,
			want:     false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
