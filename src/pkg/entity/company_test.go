package entity

import (
	"testing"
)

func TestNewCompany(t *testing.T) {
	tests := []struct {
		in   string
		want Company
	}{
		{
			in: "GREAT COMPANY",
			want: Company{
				ID:   0,
				Name: "GREAT COMPANY",
			},
		},
	}

	for _, tt := range tests {
		got := NewCompany(tt.in)
		if tt.want != *got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}
}

func TestCompanyValidate(t *testing.T) {
	tests := []struct {
		company *Company
	}{
		{&Company{
			ID:   0,
			Name: "GREAT COMPANY",
		}},
	}

	for _, tt := range tests {
		err := tt.company.Validate()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestFailedCompanyValidate(t *testing.T) {
	tests := []struct {
		company *Company
	}{
		{&Company{
			ID:   0,
			Name: "",
		}},
	}

	for _, tt := range tests {
		err := tt.company.Validate()
		if err == nil {
			t.Fatalf("Expect Error")
		}
		t.Log(err)
	}
}
