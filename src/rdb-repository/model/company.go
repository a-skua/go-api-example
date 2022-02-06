package model

import (
	"api.example.com/pkg/entity"
	"context"
	"database/sql"
	"time"
)

type Company struct {
	ID        entity.CompanyID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCompany(entity *entity.Company) *Company {
	return &Company{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func (c *Company) Entity() *entity.Company {
	return &entity.Company{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c *Company) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into companies(name, created_at, updated_at) value (?, ?, ?)",
		c.Name,
		now,
		now,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.ID = entity.CompanyID(id)
	c.CreatedAt = now
	c.UpdatedAt = now
	return nil
}

type CompanyRole struct {
	ID        int
	CompanyID entity.CompanyID
	RoleID    entity.RoleID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCompanyRole(cID entity.CompanyID, rID entity.RoleID) *CompanyRole {
	return &CompanyRole{
		CompanyID: cID,
		RoleID:    rID,
	}
}

func FindCompanyRoleTx(tx *sql.Tx, cID entity.CompanyID, rID entity.RoleID) (*CompanyRole, error) {
	cr := NewCompanyRole(cID, rID)
	err := tx.QueryRowContext(
		context.Background(),
		"select id, created_at, updated_at from company_roles where company_id=? and role_id=?",
		cID,
		rID,
	).Scan(&cr.ID, &cr.CreatedAt, &cr.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (cr *CompanyRole) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into company_roles(company_id, role_id, created_at, updated_at) value (?, ?, ?, ?)",
		cr.CompanyID,
		cr.RoleID,
		now,
		now,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	cr.ID = int(id)
	cr.CreatedAt = now
	cr.UpdatedAt = now
	return nil
}

type CompanyEmployee struct {
	ID        int
	CompanyID entity.CompanyID
	UserID    entity.UserID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCompanyEmployee(cID entity.CompanyID, uID entity.UserID) *CompanyEmployee {
	return &CompanyEmployee{
		CompanyID: cID,
		UserID:    uID,
	}
}

func FindCompanyEmployeeTx(tx *sql.Tx, cID entity.CompanyID, uID entity.UserID) (*CompanyEmployee, error) {
	ce := NewCompanyEmployee(cID, uID)
	err := tx.QueryRowContext(
		context.Background(),
		"select id, created_at, updated_at from company_employees where company_id=? and user_id=?",
		cID,
		uID,
	).Scan(&ce.ID, &ce.CreatedAt, &ce.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return ce, nil
}

func (ce *CompanyEmployee) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into company_employees(company_id, user_id, created_at, updated_at) value (?, ?, ?, ?)",
		ce.CompanyID,
		ce.UserID,
		now,
		now,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	ce.ID = int(id)
	ce.CreatedAt = now
	ce.UpdatedAt = now
	return nil
}

type EmployeeRole struct {
	ID                int
	CompanyEmployeeID int
	CompanyRoleID     int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewEmployeeRole(companyEmployeeID, companyRoleID int) *EmployeeRole {
	return &EmployeeRole{
		CompanyEmployeeID: companyEmployeeID,
		CompanyRoleID:     companyRoleID,
	}
}

func (er *EmployeeRole) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into employee_roles(company_employee_id, company_role_id, created_at, updated_at) value (?, ?, ?, ?)",
		er.CompanyEmployeeID,
		er.CompanyRoleID,
		now,
		now,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	er.ID = int(id)
	er.CreatedAt = now
	er.UpdatedAt = now
	return nil
}
