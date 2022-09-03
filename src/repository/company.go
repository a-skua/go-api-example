package repository

import (
	companies "api.example.com/pkg/company"
	"api.example.com/repository/model"
	"fmt"
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
