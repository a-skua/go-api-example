package company

import (
	"fmt"
	"reflect"
	"testing"
)

// mock
type repository struct {
	create, read, update, delete bool
	id                           ID
	company                      *Company
	err                          error
}

func (r *repository) CompanyCreate(*Company) (*Company, error) {
	if r.create {
		return r.company, r.err
	} else {
		return nil, fmt.Errorf("failed create")
	}
}

func (r *repository) CompanyRead(ID) (*Company, error) {
	if r.read {
		return r.company, r.err
	} else {
		return nil, fmt.Errorf("failed read")
	}
}

func (r *repository) CompanyUpdate(*Company) (*Company, error) {
	if r.update {
		return r.company, r.err
	} else {
		return nil, fmt.Errorf("failed update")
	}
}

func (r *repository) CompanyDelete(ID) error {
	if r.delete {
		return r.err
	} else {
		return fmt.Errorf("failed delete")
	}
}

// test
func TestNewServer(t *testing.T) {
	type test struct {
		testcase   string
		repository Repository
		want       Server
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := NewServer(tt.repository)

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			repository: new(repository),
			want:       &server{new(repository)},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Create(t *testing.T) {
	type args struct {
		company *Company
	}
	type test struct {
		testcase string
		args
		server  *server
		want    *Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Create(tt.company)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-erorr=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				company: New("GREATE COMPANY"),
			},
			server: &server{&repository{
				create: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY",
				},
				err: nil,
			}},
			want: &Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
		},
		{
			testcase: "invalid company",
			args: args{
				company: New(""),
			},
			server: &server{&repository{
				create: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY",
				},
				err: nil,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			testcase: "repository error",
			args: args{
				company: New("GREATE COMPANY"),
			},
			server: &server{&repository{
				create:  true,
				company: nil,
				err:     fmt.Errorf("internal server error"),
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Read(t *testing.T) {
	type args struct {
		id ID
	}

	type test struct {
		testcase string
		args
		server  *server
		want    *Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Read(tt.id)
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
			args: args{
				id: 1,
			},
			server: &server{&repository{
				read: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY",
				},
				err: nil,
			}},
			want: &Company{
				ID:   1,
				Name: "GREATE COMPANY",
			},
			wantErr: false,
		},
		{
			testcase: "invalid id",
			args: args{
				id: 0,
			},
			server: &server{&repository{
				read: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY",
				},
				err: nil,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			testcase: "repository error",
			args: args{
				id: 1,
			},
			server: &server{&repository{
				read:    true,
				company: nil,
				err:     fmt.Errorf("internal server error"),
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Update(t *testing.T) {
	type args struct {
		company *Company
	}

	type test struct {
		testcase string
		args
		server  *server
		want    *Company
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got, err := tt.server.Update(tt.company)
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
			args: args{
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY!!",
				},
			},
			server: &server{&repository{
				update: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY!!",
				},
				err: nil,
			}},
			want: &Company{
				ID:   1,
				Name: "GREATE COMPANY!!",
			},
			wantErr: false,
		},
		{
			testcase: "invalid id",
			args: args{
				company: &Company{
					ID:   0,
					Name: "GREATE COMPANY!!",
				},
			},
			server: &server{&repository{
				update: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY!!",
				},
				err: nil,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			testcase: "invalid company",
			args: args{
				company: &Company{
					ID:   1,
					Name: "",
				},
			},
			server: &server{&repository{
				update: true,
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY!!",
				},
				err: nil,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			testcase: "repository error",
			args: args{
				company: &Company{
					ID:   1,
					Name: "GREATE COMPANY!!",
				},
			},
			server: &server{&repository{
				update:  true,
				company: nil,
				err:     fmt.Errorf("internal server error"),
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestServer_Delete(t *testing.T) {
	type args struct {
		id ID
	}

	type test struct {
		testcase string
		args
		server  *server
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			err := tt.server.Delete(tt.id)
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}
		})
	}

	tests := []*test{
		{
			args: args{
				id: 1,
			},
			server: &server{&repository{
				delete: true,
				err:    nil,
			}},
			wantErr: false,
		},
		{
			testcase: "invalid id",
			args: args{
				id: 0,
			},
			server: &server{&repository{
				delete: true,
				err:    nil,
			}},
			wantErr: true,
		},
		{
			testcase: "repository error",
			args: args{
				id: 1,
			},
			server: &server{&repository{
				delete: true,
				err:    fmt.Errorf("internal server error"),
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
