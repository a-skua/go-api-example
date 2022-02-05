package model

import (
	"api.example.com/pkg/entity"
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        entity.UserID
	Name      string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(entity *entity.User) *User {
	return &User{
		ID:       entity.ID,
		Name:     entity.Name,
		Password: entity.Password.Hash(),
	}
}

func (u *User) Entity() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: entity.PasswordFromHash(u.Password),
	}
}

func (u *User) Create(tx *sql.Tx) error {
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
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = entity.UserID(id)
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *User) Read(db *sql.DB) error {
	err := db.QueryRowContext(
		context.Background(),
		"select name, password, created_at, updated_at from users where id = ?",
		u.ID,
	).Scan(&u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(tx *sql.Tx) error {
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
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return errors.New("rows affected not 1")
	}

	u.UpdatedAt = now
	return nil
}

func (u *User) Delete(tx *sql.Tx) error {
	result, err := tx.ExecContext(
		context.Background(),
		"delete from users where id=?",
		u.ID,
	)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return errors.New("rows affected not 1")
	}

	return nil
}
