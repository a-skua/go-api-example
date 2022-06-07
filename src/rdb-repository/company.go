package repository

import (
	"api.example.com/pkg/company"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

func CompanyCreate(db *sql.DB, company *company.Company) (*company.Company, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.CompanyCreate: %w", err)
	}

	model, err := companyCreate(tx, model.NewCompany(company))
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("rdb-repository.CompanyCreate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.CompanyCreate: %w", err)
	}

	return model.NewEntity(), nil
}

func companyCreate(tx *sql.Tx, company model.Company) (model.Company, error) {
	err := company.Create(tx)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func CompanyRead(db *sql.DB, id company.ID) (*company.Company, error) {
	model, err := companyRead(db, model.NewCompanyFromID(id))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.CompanyRead: %w", err)
	}

	return model.NewEntity(), nil
}

func companyRead(db *sql.DB, company model.Company) (model.Company, error) {
	err := company.Read(db)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func CompanyUpdate(db *sql.DB, company *company.Company) (*company.Company, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.CompanyUpdate: %w", err)
	}

	model, err := companyUpdate(tx, model.NewCompany(company))
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("rdb-repository.CompanyUpdate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.CompanyUpdate: %w", err)
	}

	return model.NewEntity(), nil
}

func companyUpdate(tx *sql.Tx, company model.Company) (model.Company, error) {
	err := company.Update(tx)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func CompanyDelete(db *sql.DB, id company.ID) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("rdb-repository.CompanyDelete: %w", err)
	}

	err = companyDelete(tx, model.NewCompanyFromID(id))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("rdb-repository.CompanyDelete: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("rdb-repository.CompanyDelete: %w", err)
	}

	return nil

}

func companyDelete(tx *sql.Tx, company model.Company) error {
	err := company.Delete(tx)
	if err != nil {
		return err
	}
	return nil
}
