package response

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		err      error
		wantErr  bool
		want
	}

	do := func(tt test) {
		t.Logf("testcase: %s", tt.testcase)

		w := httptest.NewRecorder()
		err := Error(w, tt.err)
		if hasErr := err != nil; tt.wantErr != hasErr {
			t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
		}

		res := w.Result()
		gotBody, _ := io.ReadAll(res.Body)
		res.Body.Close()

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
	}

	tests := []test{
		{
			err:     fmt.Errorf("error test"),
			wantErr: false,
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
