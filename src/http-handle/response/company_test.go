package response

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"api.example.com/pkg/company"
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

func TestCompanyRead(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		company  *company.Company
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := CompanyRead(w, tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
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
		})
	}

	tests := []*test{
		{
			testcase: "OK1",
			company: &company.Company{
				ID:        1,
				Name:      "testCompany",
				OwnerID:   1,
				UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":1,"name":"testCompany","owner_id":1,"updated_at":"2022-09-03T12:34:56Z"}}` + "\n"),
			},
		},
		{
			testcase: "OK2",
			company: &company.Company{
				ID:        2,
				Name:      "greatCompany",
				OwnerID:   2,
				UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":2,"name":"greatCompany","owner_id":2,"updated_at":"2022-09-03T12:34:56Z"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
