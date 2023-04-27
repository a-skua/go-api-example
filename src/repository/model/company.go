package model

import (
	"context"
	"fmt"

	companies "api.example.com/pkg/company"
)

type Company interface {
	Create(DB) error
	Read(DB) error
	NewEntity() *companies.Company
}

// impl Company
type company struct {
	id        companies.ID
	name      companies.Name
	createdAt dateTime
	updatedAt dateTime
	// TODO OwnerID: テーブル設計を見直すこと
}

func NewCompany(c *companies.Company) Company {
	return &company{
		name: c.Name,
	}
}

func NewCompanyFromID(id companies.ID) Company {
	return &company{
		id: id,
	}
}

func (c *company) Create(tx DB) error {
	now := currentTime()
	result, err := tx.ExecContext(
		context.TODO(),
		"insert into `companies`(`name`, `created_at`, `updated_at`) value (?, ?, ?)",
		c.name,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("repository/model.Company.Create: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("repository/model.Company.Create: %w", err)
	}

	c.id = companies.ID(id)
	c.createdAt = now
	c.updatedAt = now
	return nil
}

func (c *company) Read(tx DB) error {
	err := tx.QueryRowContext(
		context.TODO(),
		"select `id`, `name`, `created_at`, `updated_at` from `companies` where `id`=?",
		c.id,
	).Scan(&c.id, &c.name, &c.createdAt, &c.updatedAt)

	if err != nil {
		return fmt.Errorf("repository/model.Company.Read: %w", err)
	}

	return nil
}

func (c *company) NewEntity() *companies.Company {
	return &companies.Company{
		ID:   c.id,
		Name: c.name,
		// TODO OwnerID
		UpdatedAt: c.updatedAt,
	}
}
