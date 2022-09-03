package handle

import (
	"api.example.com/pkg/company"
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// mock
type companyServer struct {
	company *company.Company
	err     error
	// flag
	create bool
	// test
	t *testing.T
}

// mock
type makeServer func(t *testing.T) company.Server

func (s *companyServer) Create(*company.Company) (*company.Company, error) {
	s.t.Helper()

	if s.create {
		return s.company, s.err
	}
	s.t.Fatal("invalid Create")
	panic("invalid Create")
}

func TestCompanyHanlder_create(t *testing.T) {
	type args struct {
		url  string
		body []byte
	}

	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}

	type test struct {
		name string
		args
		makeServer makeServer
		want       want
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer(tt.body))
			w := httptest.NewRecorder()

			s := newServices()
			s.Company = tt.makeServer(t)

			New(s).ServeHTTP(w, r)

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
			args: args{
				url:  "http://api.example.com/company",
				body: []byte(`{"company":{"name":"GREATE COMPANY","owner_id":1}}`),
			},
			makeServer: func(t *testing.T) company.Server {
				return &companyServer{
					company: &company.Company{
						ID:        2,
						Name:      "GREATE COMPANY",
						OwnerID:   1,
						UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
					},
					err:    nil,
					create: true,
					t:      t,
				}
			},
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        []byte(`{"company":{"id":2,"name":"GREATE COMPANY","owner_id":1,"updated_at":"2022-09-03T12:34:56Z"}}` + "\n"),
			},
		},
		{
			name: "invalid request",
			args: args{
				url:  "http://api.example.com/company",
				body: []byte(``),
			},
			makeServer: func(t *testing.T) company.Server {
				return &companyServer{
					t: t,
				}
			},
			want: want{
				statusCode:  500,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
		{
			name: "failed create",
			args: args{
				url:  "http://api.example.com/company",
				body: []byte(`{"company":{"name":"GREATE COMPANY","owner_id":1}}`),
			},
			makeServer: func(t *testing.T) company.Server {
				return &companyServer{
					err:    errors.New("internal server error"),
					create: true,
					t:      t,
				}
			},
			want: want{
				statusCode:  500,
				contentType: "application/json",
				body:        []byte(`{"error":{}}` + "\n"),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
