package repository

import (
	companies "api.example.com/pkg/company"
	"api.example.com/repository/model"
	"errors"
	"reflect"
	"testing"
	"time"
)

type makeModelCompany func(*testing.T) model.Company

// mock
type modelCompany struct {
	entity *companies.Company
	err    error
	// flags
	create, newEntity bool
	// test
	t *testing.T
}

func (c *modelCompany) Create(tx model.DB) error {
	c.t.Helper()
	if c.create {
		return c.err
	}

	c.t.Fatal("invalid Create")
	panic("invalid Create")
}

func (c *modelCompany) NewEntity() *companies.Company {
	c.t.Helper()
	if c.newEntity {
		return c.entity
	}
	c.t.Fatal("invalid NewEntity")
	panic("invalid NewEntity")
}

func TestCompanyCreate(t *testing.T) {
	t.Skip("TODO OwnerID")

	type test struct {
		name        string
		tx          Transaction
		makeCompany makeModelCompany
		want        *companies.Company
		wantErr     bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := companyCreate(tt.tx, tt.makeCompany(t))
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
			tx: &transaction{
				commit: true,
			},
			makeCompany: func(t *testing.T) model.Company {
				return &modelCompany{
					entity: &companies.Company{
						ID:        1,
						Name:      "GREATE COMPANY",
						OwnerID:   2,
						UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
					},
					create:    true,
					newEntity: true,
					t:         t,
				}
			},
			want: &companies.Company{
				ID:        1,
				Name:      "GREATE COMPANY",
				OwnerID:   2,
				UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "failed create",
			tx: &transaction{
				rollback: true,
			},
			makeCompany: func(t *testing.T) model.Company {
				return &modelCompany{
					err:    errors.New("test error"),
					create: true,
					t:      t,
				}
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed commit",
			tx: &transaction{
				errCommit: errors.New("test error"),
				commit:    true,
			},
			makeCompany: func(t *testing.T) model.Company {
				return &modelCompany{
					create: true,
					t:      t,
				}
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
