package model

import (
	companies "api.example.com/pkg/company"
	"context"
	"fmt"
)

type Company interface {
	Create(DB) error
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

func (c *company) NewEntity() *companies.Company {
	return &companies.Company{
		ID:   c.id,
		Name: c.name,
		// TODO OwnerID
		UpdatedAt: c.updatedAt,
	}
}
