package response

import (
	"api.example.com/pkg/company"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestWriteCompany(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		name    string
		company *company.Company
		want    want
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := WriteCompany(w, tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			got := w.Result()
			defer got.Body.Close()

			gotBody, _ := io.ReadAll(got.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("body want=%s, got=%s.", tt.want.body, gotBody)
			}

			gotContentType := got.Header.Get("Content-Type")
			if tt.want.contentType != gotContentType {
				t.Fatalf("Content-Type want=%v, got=%v.", tt.want.contentType, gotContentType)
			}

			gotStatusCode := got.StatusCode
			if tt.want.statusCode != gotStatusCode {
				t.Fatalf("Status-Code want=%v, got=%v.", tt.want.statusCode, gotStatusCode)
			}
		})
	}

	tests := []*test{
		{
			name: "ok",
			company: &company.Company{
				ID:        2,
				Name:      "GREATE COMPANY",
				OwnerID:   1,
				UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":2,"name":"GREATE COMPANY","owner_id":1,"updated_at":"2022-09-03T12:34:56Z"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
