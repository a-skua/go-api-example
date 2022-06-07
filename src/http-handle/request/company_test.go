package request

import (
	"api.example.com/pkg/company"
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCompanyCreate(t *testing.T) {
	type test struct {
		testcase string
		body     []byte
		password string
		want     *company.Company
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest("POST", "http://api.example.com/company", bytes.NewBuffer(tt.body))

			got, err := CompanyCreate(r)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			body:    []byte(`{"company":{"name":"GREATE COMPANY"}}`),
			want:    company.New("GREATE COMPANY"),
			wantErr: false,
		},
		{
			testcase: "invalid json",
			body:     []byte(""),
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompanyRead(t *testing.T) {
	type test struct {
		testcase string
		url      string
		want     company.ID
		wantErr  bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)

			var (
				got company.ID
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/company/{company_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = CompanyRead(r)
			})
			router.ServeHTTP(nil, r)

			if hasErr := err != nil; tt.wantErr != hasErr {
				t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			url:     "http://api.example.com/company/1",
			want:    1,
			wantErr: false,
		},
		{
			url:     "http://api.example.com/company/foo",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompanyUpdate(t *testing.T) {
	type test struct {
		testcase string
		url      string
		body     []byte
		want     *company.Company
		wantErr  bool
	}

	do := func(tt test) {
		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest("PUT", tt.url, bytes.NewBuffer(tt.body))

			var (
				got *company.Company
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/company/{company_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = CompanyUpdate(r)
			})
			router.ServeHTTP(nil, r)
			if hasErr := err != nil; tt.wantErr != hasErr {
				t.Fatalf("want-err=%v, err=%v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []test{
		{
			url:  "http://api.example.com/company/1",
			body: []byte(`{"company":{"name":"GREATE COMPANY"}}`),
			want: &company.Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
		},
		{
			testcase: "empty body",
			url:      "http://api.example.com/company/1",
			body:     []byte(""),
			want:     nil,
			wantErr:  true,
		},
		{
			testcase: "invalid company_id",
			url:      "http://api.example.com/company/xxx",
			body:     []byte(`{"company":{"name":"GREATE COMPANY"}}`),
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompanyDelete(t *testing.T) {
	type test struct {
		testcase string
		url      string
		want     company.ID
		wantErr  bool
	}

	do := func(tt test) {
		t.Logf("testcace: %v", tt.testcase)

		t.Run(tt.testcase, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)

			var (
				got company.ID
				err error
			)
			router := mux.NewRouter()
			router.HandleFunc("/company/{company_id}", func(w http.ResponseWriter, r *http.Request) {
				got, err = CompanyDelete(r)
			})
			router.ServeHTTP(nil, r)

			if tt.wantErr != (err != nil) {
				t.Fatalf("want-err=%v, err=%v.", tt.wantErr, err)
			}

			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []test{
		{
			url:     "http://api.example.com/company/2",
			want:    2,
			wantErr: false,
		},
		{
			testcase: "invalid company_id",
			url:      "http://api.example.com/company/xxx",
			want:     0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
