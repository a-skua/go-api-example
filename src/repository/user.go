package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/repository/model"
	"fmt"
)

func UserCreate(tx Transaction, model model.User) (*users.User, error) {
	err := model.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository.UserCreate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("repository.UserCreate: %w", err)
	}

	return model.NewEntity(), nil
}

func UserRead(db DB, model model.User) (*users.User, error) {
	err := model.Read(db)
	if err != nil {
		return nil, fmt.Errorf("repository.UserRead: %w", err)
	}

	return model.NewEntity(), nil
}

func UserUpdate(tx Transaction, model model.User) (*users.User, error) {
	err := model.Update(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository.UserUpdate: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("repository.UserUpdate: %w", err)
	}

	return model.NewEntity(), nil
}

func UserDelete(tx Transaction, model model.User) error {
	err := model.Delete(tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repository.UserDelete: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("repository.UserDelete: %w", err)
	}

	return nil
}
