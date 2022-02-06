package model

import (
	"api.example.com/pkg/entity"
	"context"
	"database/sql"
	"time"
)

type Role struct {
	ID        entity.RoleID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRole(entity *entity.Role) *Role {
	return &Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func (r *Role) Entity(c *entity.Company) *entity.Role {
	return &entity.Role{
		Company: c,
		ID:      r.ID,
		Name:    r.Name,
	}
}

func (r *Role) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into roles(name, created_at, updated_at) value (?, ?, ?)",
		r.Name,
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

	r.ID = entity.RoleID(id)
	r.CreatedAt = now
	r.UpdatedAt = now
	return nil
}
