package company

import (
	"reflect"
	"testing"
)

func TestID_valid(t *testing.T) {
	type test struct {
		testcase string
		id       ID
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.id.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			id:   0,
			want: false,
		},
		{
			id:   1,
			want: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestName_valid(t *testing.T) {
	type test struct {
		testcase string
		name     Name
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.name.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "",
			want: false,
		},
		{
			name: "1",
			want: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		name Name
	}

	type test struct {
		testcase string
		args
		want *Company
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := New(tt.name)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				name: "Bob",
			},
			want: &Company{
				Name: "Bob",
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompany_valid(t *testing.T) {
	type test struct {
		testcase string
		company  *Company
		want     bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := tt.company.valid()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			company: &Company{
				Name: "Bob",
			},
			want: true,
		},
		{
			testcase: "invalid name",
			company: &Company{
				Name: "",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
