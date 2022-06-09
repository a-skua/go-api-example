package repository

import (
	"api.example.com/pkg/company"
	users "api.example.com/pkg/user"
	"database/sql"
)

type Repository interface {
	users.Repository
	company.Repository
	Close() error
}

type rdb struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &rdb{
		db: db,
	}
}

func (r *rdb) Close() error {
	return r.db.Close()
}

func (r *rdb) UserCreate(u *users.User) (*users.User, error) {
	return UserCreate(r.db, u)
}

func (r *rdb) UserRead(id users.ID) (*users.User, error) {
	return UserRead(r.db, id)
}

func (r *rdb) UserUpdate(u *users.User) (*users.User, error) {
	return UserUpdate(r.db, u)
}

func (r *rdb) UserDelete(id users.ID) error {
	return UserDelete(r.db, id)
}

func (r *rdb) CompanyCreate(c *company.Company) (*company.Company, error) {
	return CompanyCreate(r.db, c)
}

func (r *rdb) CompanyRead(id company.ID) (*company.Company, error) {
	return CompanyRead(r.db, id)
}

func (r *rdb) CompanyUpdate(c *company.Company) (*company.Company, error) {
	return CompanyUpdate(r.db, c)
}

func (r *rdb) CompanyDelete(id company.ID) error {
	return CompanyDelete(r.db, id)
}
