package repository

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

type rdb struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Repository {
	return &rdb{db}
}

func (r *rdb) Begin() repository.Tx {
	return newTx(r.db.Begin())
}

func (r *rdb) UserCreate(entity *entity.User) (*entity.User, error) {
	user := model.NewUser(entity)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository UserCreate: %w", err)
	}

	err = user.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository UserCreate: %w", err)
	}

	tx.Commit()
	return user.Entity(), nil
}

func (r *rdb) UserRead(userid entity.UserID) (*entity.User, error) {
	user := &model.User{ID: userid}

	err := user.Read(r.db)
	if err != nil {
		return nil, fmt.Errorf("repository UserRead: %w", err)
	}

	return user.Entity(), nil
}

func (r *rdb) UserUpdate(entity *entity.User) (*entity.User, error) {
	user := model.NewUser(entity)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository UserUpdate: %w", err)
	}

	err = user.Update(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository UserUpdate: %w", err)
	}

	tx.Commit()
	return user.Entity(), nil
}

func (r *rdb) UserDelete(userid entity.UserID) error {
	user := &model.User{ID: userid}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("repository UserDelete: %w", err)
	}

	err = user.Delete(tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repository UserDelete: %w", err)
	}

	tx.Commit()
	return nil
}

func (r *rdb) CompanyCreateTx(entity *entity.Company, trans repository.Tx) (*entity.Company, repository.Tx) {
	company := model.NewCompany(entity)
	tx, err := castTx(trans)
	if err != nil {
		return nil, tx.newState(err)
	}

	err = company.Create(tx.sql)
	if err != nil {
		return nil, tx.newState(err)
	}

	return company.Entity(), tx.newState(nil)
}

func (r *rdb) CompanyAddEmployeeTx(cID entity.CompanyID, uID entity.UserID, trans repository.Tx) repository.Tx {
	ce := model.NewCompanyEmployee(cID, uID)
	tx, err := castTx(trans)
	if err != nil {
		return tx.newState(err)
	}

	err = ce.Create(tx.sql)
	if err != nil {
		return tx.newState(err)
	}

	return tx.newState(nil)
}

func (r *rdb) RoleCreateTx(entity *entity.Role, trans repository.Tx) (*entity.Role, repository.Tx) {
	company := entity.Company
	role := model.NewRole(entity)
	tx, err := castTx(trans)
	if err != nil {
		return nil, tx.newState(err)
	}

	err = role.Create(tx.sql)
	if err != nil {
		return nil, tx.newState(err)
	}

	cr := model.NewCompanyRole(company.ID, role.ID)
	err = cr.Create(tx.sql)
	if err != nil {
		return nil, tx.newState(err)
	}

	return role.Entity(company), tx.newState(nil)
}

func (r *rdb) EmployeeAddRoleTx(cID entity.CompanyID, uID entity.UserID, rID entity.RoleID, trans repository.Tx) repository.Tx {
	tx, err := castTx(trans)
	if err != nil {
		return tx.newState(err)
	}

	ce, err := model.FindCompanyEmployeeTx(tx.sql, cID, uID)
	if err != nil {
		return tx.newState(err)
	}

	cr, err := model.FindCompanyRoleTx(tx.sql, cID, rID)
	if err != nil {
		return tx.newState(err)
	}

	er := model.NewEmployeeRole(ce.ID, cr.ID)
	err = er.Create(tx.sql)
	if err != nil {
		return tx.newState(err)
	}

	return tx.newState(nil)
}
