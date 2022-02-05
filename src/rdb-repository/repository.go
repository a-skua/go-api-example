package repository

import (
	"api.example.com/pkg/entity"
	"api.example.com/pkg/repository"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

type Repository interface {
	repository.User
}

type rdb struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &rdb{db}
}

func (r *rdb) UserCreate(entity *entity.User) (*entity.User, error) {
	user := model.NewUser(entity)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("repository UserCreate: %w", err)
	}

	err = user.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("repository UserCreate: %w", err)
	}

	tx.Commit()
	return user.Entity(), nil
}
