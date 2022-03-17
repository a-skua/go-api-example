package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

func userCreate(db *sql.DB, u model.User) (*users.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.user.create: %w", err)
	}

	err = u.Create(tx)
	if err != nil {
		tx.Rollback() // TODO error handling
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.user.create: %w", err)
	}
	return u.NewEntity(), nil
}

func userRead(db *sql.DB, u model.User) (*users.User, error) {
	err := u.Read(db)
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.user.read: %w", err)
	}

	return u.NewEntity(), nil
}

func userUpdate(db *sql.DB, u model.User) (*users.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.user.update: %w", err)
	}

	err = u.Update(tx)
	if err != nil {
		tx.Rollback() // TODO error handling
		return nil, fmt.Errorf("rdb-repository.user.update: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.user.update: %w", err)
	}
	return u.NewEntity(), nil
}

func userDelete(db *sql.DB, u model.User) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("rdb-repository.user.delete: %w", err)
	}

	err = u.Delete(tx)
	if err != nil {
		tx.Rollback() // TODO error handling
		return fmt.Errorf("rdb-repository.user.delete: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("rdb-repository.user.delete: %w", err)
	}
	return nil
}
