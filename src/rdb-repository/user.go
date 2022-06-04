package repository

import (
	"api.example.com/pkg/user"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

// UserCreate
func UserCreate(db *sql.DB, user *user.User) (*user.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserCreate: %w", err)
	}

	model, err := userCreate(tx, model.NewUser(user))
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("rdb-repository.UserCreate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserCreate: %w", err)
	}

	return model.NewEntity(), nil
}

func userCreate(tx *sql.Tx, user model.User) (model.User, error) {
	err := user.Create(tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserRead
func UserRead(db *sql.DB, id user.ID) (*user.User, error) {
	model, err := userRead(db, model.NewUserFromID(id))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserRead: %w", err)
	}

	return model.NewEntity(), nil
}

func userRead(db *sql.DB, user model.User) (model.User, error) {
	err := user.Read(db)
	if err != nil {
		return nil, nil
	}

	return user, nil
}

// UserUpdate
func UserUpdate(db *sql.DB, user *user.User) (*user.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserUpdate: %w", err)
	}

	model, err := userUpdate(tx, model.NewUser(user))
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("rdb-repository.UserUpdate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserUpdate: %w", err)
	}

	return model.NewEntity(), nil
}

func userUpdate(tx *sql.Tx, user model.User) (model.User, error) {
	err := user.Update(tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserDelete
func UserDelete(db *sql.DB, id user.ID) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("rdb-repository.UserDelete: %w", err)
	}

	err = userDelete(tx, model.NewUserFromID(id))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("rdb-repository.UserDelete: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("rdb-repository.UserDelete: %w", err)
	}

	return nil
}

func userDelete(tx *sql.Tx, user model.User) error {
	err := user.Delete(tx)
	if err != nil {
		return err
	}

	return nil
}
