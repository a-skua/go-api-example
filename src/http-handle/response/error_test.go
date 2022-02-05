package response

import (
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		in   int
		want []byte
	}{
		{
			400,
			[]byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		err := Error(w, tt.in)
		if err != nil {
			t.Fatal(err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}
}
