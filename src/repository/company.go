package repository

import (
	"fmt"

	companies "api.example.com/pkg/company"
	"api.example.com/repository/model"
)

func companyCreate(tx Transaction, model model.Company) (*companies.Company, error) {
	err := model.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository.CompanyCreate: %w:", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("repository.CompanyCreate: %w:", err)
	}

	return model.NewEntity(), nil
}

func companyRead(db DB, model model.Company) (*companies.Company, error) {
	err := model.Read(db)
	if err != nil {
		return nil, fmt.Errorf("repository.CompanyRead: %w:", err)
	}

	return model.NewEntity(), nil
}
