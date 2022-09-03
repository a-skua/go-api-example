package company

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name    Name
		ownerid OwnerID
	}

	type test struct {
		name string
		args
		want *Company
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.name, tt.ownerid)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			args: args{
				name:    "foo",
				ownerid: 1,
			},
			want: &Company{
				Name:    "foo",
				OwnerID: 1,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompany_validCreate(t *testing.T) {
	type test struct {
		name    string
		company *Company
		want    bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.company.validCreate()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:    "ok",
			company: New("GREATE COMPANY", 1),
			want:    true,
		},
		{
			name:    "name valid min",
			company: New("G", 1),
			want:    true,
		},
		{
			name:    "name valid max",
			company: New(Name(strings.Repeat("1", 255)), 1),
			want:    true,
		},
		{
			name:    "name invalid min",
			company: New("", 1),
			want:    false,
		},
		{
			name:    "name invalid max",
			company: New(Name(strings.Repeat("1", 256)), 1),
			want:    false,
		},
		{
			name:    "name invalid ownerid",
			company: New("GREATE COMPANY", 0),
			want:    false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
