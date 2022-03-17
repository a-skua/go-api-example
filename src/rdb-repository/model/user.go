package model

import (
	users "api.example.com/pkg/user"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User interface {
	Create(*sql.Tx) error
	Read(*sql.DB) error
	Update(*sql.Tx) error
	Delete(*sql.Tx) error
	NewEntity() *users.User
}

type user struct {
	ID        users.ID
	Name      string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(u *users.User) User {
	return &user{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password.Hash(),
	}
}

func NewUserFromID(id users.ID) User {
	return &user{
		ID: id,
	}
}

func (u *user) NewEntity() *users.User {
	return &users.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: users.NewPasswordFromHash(u.Password),
	}
}

func (u *user) Create(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"insert into users(name, password, created_at, updated_at) value (?, ?, ?, ?)",
		u.Name,
		u.Password,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Create: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Create: %w", err)
	}

	u.ID = users.ID(id)
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *user) Read(db *sql.DB) error {
	err := db.QueryRowContext(
		context.Background(),
		"select name, password, created_at, updated_at from users where id = ?",
		u.ID,
	).Scan(&u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Read: %w", err)
	}
	return nil
}

func (u *user) Update(tx *sql.Tx) error {
	now := time.Now()
	result, err := tx.ExecContext(
		context.Background(),
		"update users set name=?, password=?, updated_at=? where id=?",
		u.Name,
		u.Password,
		now,
		u.ID,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Update: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Update: %w", err)
	}

	if count != 1 {
		return fmt.Errorf("rdb-repository/model.User.Update: rows affected not 1 (affected=%d)", count)
	}

	u.UpdatedAt = now
	return nil
}

func (u *user) Delete(tx *sql.Tx) error {
	result, err := tx.ExecContext(
		context.Background(),
		"delete from users where id=?",
		u.ID,
	)
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Delete: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rdb-repository/model.User.Delete: %w", err)
	}

	if count != 1 {
		return fmt.Errorf("rdb-repository/model.User.Delete: rows affected not 1 (affected=%d)", count)
	}

	return nil
}
