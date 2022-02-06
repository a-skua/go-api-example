package entity

import (
	"reflect"
	"testing"
)

func TestNewRoleAdmin(t *testing.T) {
	tests := []struct {
		in   *Company
		want *Role
	}{
		{
			in: &Company{
				ID:   1,
				Name: "GREAT COMPANY",
			},
			want: &Role{
				Company: &Company{
					ID:   1,
					Name: "GREAT COMPANY",
				},
				ID:   0,
				Name: "管理者",
			},
		},
	}

	for _, tt := range tests {
		got := NewRoleAdmin(tt.in)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}
}
