package request

import (
	"api.example.com/pkg/company"
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewCompanyCreate(t *testing.T) {
	type test struct {
		name    string
		url     string
		body    []byte
		want    *company.Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer(tt.body))

			got, err := NewCompanyCreate(r)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "ok",
			url:  "http://api.example.com/company",
			body: []byte(`{
  "company": {
    "name": "GREATE COMPANY",
    "owner_id": 1
  }
}`),
			want:    company.New("GREATE COMPANY", 1),
			wantErr: false,
		},
		{
			name:    "invalid request",
			url:     "http://api.example.com/company",
			body:    []byte{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
