package response

import (
	"api.example.com/pkg/entity"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

// mock Password
type plainPassword string

func (pw plainPassword) Verify(plain string) bool {
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

func TestUser(t *testing.T) {
	tests := []struct {
		in   *entity.User
		want []byte
	}{
		{
			&entity.User{
				ID:        1,
				Name:      "Bob",
				Password:  plainPassword("password"),
				Companies: []*entity.Company{},
			},
			[]byte(`{"user":{"id":1,"name":"Bob","password":"*****","companies":[]}}` + "\n"),
		},
		{
			&entity.User{
				ID:       2,
				Name:     "Alice",
				Password: plainPassword("password"),
				Companies: []*entity.Company{
					{ID: 1, Name: "GREATE COMPANY"},
				},
			},
			[]byte(`{"user":{"id":2,"name":"Alice","password":"*****","companies":[{"id":1,"name":"GREATE COMPANY"}]}}` + "\n"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		err := user(w, tt.in)
		if err != nil {
			t.Fatal(err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}
}
