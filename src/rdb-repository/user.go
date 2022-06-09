package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

// UserCreate
func UserCreate(db *sql.DB, user *users.User) (*users.User, error) {
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

func userCreate(tx model.DB, user model.User) (model.User, error) {
	err := user.Create(tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserRead
func UserRead(db *sql.DB, id users.ID) (*users.User, error) {
	model, err := userRead(db, model.NewUserFromID(id))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.UserRead: %w", err)
	}

	return model.NewEntity(), nil
}

func userRead(tx model.DB, user model.User) (model.User, error) {
	err := user.Read(tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserUpdate
func UserUpdate(db *sql.DB, user *users.User) (*users.User, error) {
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

func userUpdate(tx model.DB, user model.User) (model.User, error) {
	err := user.Update(tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserDelete
func UserDelete(db *sql.DB, id users.ID) error {
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

func userDelete(tx model.DB, user model.User) error {
	err := user.Delete(tx)
	if err != nil {
		return err
	}

	return nil
}
