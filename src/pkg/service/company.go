package service

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"fmt"
)

type Company struct {
	repository repository.Company
}

func NewCompany(r repository.Company) *Company {
	return &Company{r}
}

func (c *Company) Create(newCompany *entity.Company, authID entity.UserID) (*entity.Company, error) {
	err := newCompany.Validate()
	if err != nil {
		return nil, err
	}

	// トランザクション開始
	tx := c.repository.Begin()
	if err := tx.Error(); err != nil {
		return nil, fmt.Errorf("service.Company.Create: %w", err)
	}

	// 会社の作成
	company, tx := c.repository.CompanyCreateTx(newCompany, tx)
	if err := tx.Error(); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 従業員の追加
	tx = c.repository.CompanyAddEmployeeTx(company.ID, authID, tx)
	if err := tx.Error(); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 監視者の作成
	role := entity.NewRoleAdmin(company)
	role, tx = c.repository.RoleCreateTx(role, tx)
	if err := tx.Error(); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 従業員を管理者に設定
	tx = c.repository.EmployeeAddRoleTx(company.ID, authID, role.ID, tx)
	if err := tx.Error(); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 確定
	tx.Commit()
	return company, nil
}
