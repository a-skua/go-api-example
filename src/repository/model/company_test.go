package model

import (
	companies "api.example.com/pkg/company"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewCompany(t *testing.T) {
	type test struct {
		name    string
		company *companies.Company
		want    Company
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCompany(tt.company)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:    "ok (create)",
			company: companies.New("GREATE COMPANY", 1),
			want: &company{
				name: "GREATE COMPANY",
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompany_Create(t *testing.T) {
	tableLock.Lock()
	defer tableLock.Unlock()

	db := newDB()
	defer db.Close()
	defer db.Exec("delete from companies")

	type test struct {
		name    string
		db      DB
		company *companies.Company
		want    *company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCompany(tt.company).(*company)

			err := got.Create(tt.db)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			wantTime := time.Now()
			tt.want.id = got.id
			testDiffTime(t, wantTime, got.createdAt)
			tt.want.createdAt = got.createdAt
			testDiffTime(t, wantTime, got.updatedAt)
			tt.want.updatedAt = got.updatedAt

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}

			// 作成されていることの確認
			err = db.
				QueryRow("select id, name, created_at, updated_at from companies where id = ?", got.id).
				Scan(&got.id, &got.name, &got.updatedAt, &got.createdAt)
			if err != nil {
				t.Fatal(err)
			}

			testDiffTime(t, wantTime, got.createdAt)
			tt.want.createdAt = got.createdAt
			testDiffTime(t, wantTime, got.updatedAt)
			tt.want.updatedAt = got.updatedAt

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:    "ok",
			db:      db,
			company: companies.New("GREATE COMPANY", 1),
			want: &company{
				name: "GREATE COMPANY",
			},
			wantErr: false,
		},
		{
			name: "failed ExecContext",
			db: &testdb{
				err:         errors.New("test error"),
				execContext: true,
			},
			company: companies.New("GREATE COMPANY", 1),
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed result.LastInsertId",
			db: &testdb{
				result: &queryResult{
					err:          errors.New("test error"),
					lastInsertID: true,
				},
				execContext: true,
			},
			company: companies.New("GREATE COMPANY", 1),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCompany_NewEntity(t *testing.T) {
	t.Skip("TODO OwnerID")

	type test struct {
		name    string
		company Company
		want    *companies.Company
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.company.NewEntity()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "ok",
			company: &company{
				id:        1,
				name:      "GREATE COMPANY",
				updatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
			want: &companies.Company{
				ID:   1,
				Name: "GREATE COMPANY",
				// TODO OwnerID
				UpdatedAt: time.Date(2022, 9, 3, 12, 34, 56, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
