package model

import (
	companies "api.example.com/pkg/company"
	"context"
	"fmt"
)

type Company interface {
	Create(DB) error
	Read(DB) error
	Update(DB) error
	Delete(DB) error
	NewEntity() *companies.Company
}

// impl Company
type company struct {
	ID        companies.ID
	Name      companies.Name
	CreatedAt dateTime
	UpdatedAt dateTime
}

func NewCompany(c *companies.Company) Company {
	return &company{
		ID:   c.ID,
		Name: c.Name,
	}
}

func NewCompanyFromID(id companies.ID) Company {
	return &company{
		ID: id,
	}
}

func (c *company) Create(tx DB) error {
	now := currentTime()
	result, err := tx.ExecContext(
		context.TODO(),
		"insert into `companies`(`name`, `created_at`, `updated_at`) value (?, ?, ?)",
		c.Name,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Create: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Create: %w", err)
	}

	c.ID = companies.ID(id)
	c.CreatedAt = now
	c.UpdatedAt = now
	return nil
}

func (c *company) Read(tx DB) error {
	err := tx.QueryRowContext(
		context.TODO(),
		"select `name`, `created_at`, `updated_at` from `companies` where `id`=?",
		c.ID,
	).Scan(&c.Name, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Read: %w", err)
	}
	return nil
}

func (c *company) Update(tx DB) error {
	now := currentTime()
	result, err := tx.ExecContext(
		context.TODO(),
		"update `companies` set `name`=?, `updated_at`=? where `id`=?",
		c.Name,
		now,
		c.ID,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Update: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Update: %w", err)
	}
	if count != 1 {
		return fmt.Errorf("rdb-repository/model.Company.Update: rows affected not 1 (affected=%d)", count)
	}

	c.UpdatedAt = now
	return nil
}

func (c *company) Delete(tx DB) error {
	result, err := tx.ExecContext(
		context.TODO(),
		"delete from `companies` where `id`=?",
		c.ID,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Delete: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.Company.Delete: %w", err)
	}

	if count != 1 {
		return fmt.Errorf("rdb-repository/model.Company.Delete: rows affected not 1 (affected=%d)", count)
	}

	return nil
}

func (c *company) NewEntity() *companies.Company {
	return &companies.Company{
		ID:   c.ID,
		Name: c.Name,
	}
}
