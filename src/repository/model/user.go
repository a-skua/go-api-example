package model

import (
	users "api.example.com/pkg/user"
	"api.example.com/pkg/user/password"
	"context"
	"fmt"
)

type passwordHash []byte

func (ph passwordHash) String() string {
	return string(ph)
}

type User interface {
	Create(DB) error
	Read(DB) error
	Update(DB) error
	Delete(DB) error
	NewEntity() *users.User
}

// impl User
type user struct {
	ID        users.ID
	Name      users.Name
	Password  passwordHash
	CreatedAt dateTime
	UpdatedAt dateTime
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
		ID:        u.ID,
		Name:      u.Name,
		Password:  password.FromHash(u.Password),
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *user) Create(tx DB) error {
	now := currentTime()
	result, err := tx.ExecContext(
		context.TODO(),
		"insert into `users`(`name`, `password`, `created_at`, `updated_at`) value (?, ?, ?, ?)",
		u.Name,
		u.Password,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("repository/model.User.Create: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("repository/model.User.Create: %w", err)
	}

	u.ID = users.ID(id)
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *user) Read(tx DB) error {
	err := tx.QueryRowContext(
		context.TODO(),
		"select `name`, `password`, `created_at`, `updated_at` from `users` where `id`=?",
		u.ID,
	).Scan(&u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return fmt.Errorf("repository/model.User.Read: %w", err)
	}
	return nil
}

func (u *user) Update(tx DB) error {
	now := currentTime()
	result, err := tx.ExecContext(
		context.TODO(),
		"update `users` set `name`=?, `password`=?, `updated_at`=? where `id`=?",
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

func (u *user) Delete(tx DB) error {
	result, err := tx.ExecContext(
		context.TODO(),
		"delete from `users` where `id`=?",
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
