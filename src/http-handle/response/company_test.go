package response

import (
	"api.example.com/pkg/company"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCompanyCreate(t *testing.T) {
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
			err := CompanyCreate(w, tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%v, got=%v.", tt.want.body, gotBody)
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
			company: &company.Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":1,"name":"GREATE COMPANY"}}` + "\n"),
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
				t.Fatalf("want=%v, got=%v.", tt.want.body, gotBody)
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
			company: &company.Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":1,"name":"GREATE COMPANY"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompanyUpdate(t *testing.T) {
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
			err := CompanyUpdate(w, tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%v, got=%v.", tt.want.body, gotBody)
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
			company: &company.Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":1,"name":"GREATE COMPANY"}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompanyDelete(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		testcase string
		wantErr  bool
		want
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := CompanyDelete(w)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			res := w.Result()
			defer res.Body.Close()

			gotBody, _ := io.ReadAll(res.Body)
			if !reflect.DeepEqual(tt.want.body, gotBody) {
				t.Fatalf("want=%v, got=%v.", tt.want.body, gotBody)
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
			wantErr: false,
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        []byte(`{"company":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
