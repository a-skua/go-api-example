package response

import (
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	type test struct {
		testcase string
		in       int
		wantErr  bool
		want     []byte
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := Error(w, tt.in)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		got, _ := io.ReadAll(res.Body)
		res.Body.Close()

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("response.user: want=%s, got=%s.", tt.want, got)
		}
	}

	tests := []test{
		{
			testcase: "success",
			in:       400,
			wantErr:  false,
			want:     []byte(`{"error":{"status":"Bad Request","status_code":400}}` + "\n"),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
