package repository

import (
	"database/sql"
	"fmt"

	companies "api.example.com/pkg/company"
	users "api.example.com/pkg/user"
	"api.example.com/repository/model"
)

type DB interface {
	model.DB
	Begin() (*sql.Tx, error)
	Close() error
}

type Transaction interface {
	model.DB
	Commit() error
	Rollback() error
}

type Repository interface {
	users.Repository
	companies.Repository
	Close() error
}

type repository struct {
	db DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Close() error {
	return r.db.Close()
}

func (r *repository) UserCreate(u *users.User) (*users.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository.UserCreate: %w", err)
	}

	return UserCreate(tx, model.NewUser(u))
}

func (r *repository) UserRead(id users.ID) (*users.User, error) {
	return UserRead(r.db, model.NewUserFromID(id))
}

func (r *repository) UserUpdate(u *users.User) (*users.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository.UserUpdate: %w", err)
	}

	return UserUpdate(tx, model.NewUser(u))
}

func (r *repository) UserDelete(id users.ID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("repository.UserDelete: %w", err)
	}

	return UserDelete(tx, model.NewUserFromID(id))
}

func (r *repository) CompanyCreate(c *companies.Company) (*companies.Company, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository.CompanyCreate: %w", err)
	}

	return companyCreate(tx, model.NewCompany(c))
}

func (r *repository) CompanyRead(id companies.ID) (*companies.Company, error) {
	return companyRead(r.db, model.NewCompanyFromID(id))
}
