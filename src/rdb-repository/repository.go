package repository

import (
	users "api.example.com/pkg/user"
	"api.example.com/rdb-repository/model"
	"database/sql"
	"fmt"
)

type Repository interface {
	users.Repository
	Close() error
}

type rdb struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &rdb{
		db: db,
	}
}

func (r *rdb) Close() error {
	return r.db.Close()
}

func (r *rdb) UserCreate(u *users.User) (*users.User, error) {
	u, err := userCreate(r.db, model.NewUser(u))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.rdb.UserCreate: %w", err)
	}
	return u, nil
}

func (r *rdb) UserRead(userid users.ID) (*users.User, error) {
	u, err := userRead(r.db, model.NewUserFromID(userid))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.rdb.UserRead: %w", err)
	}
	return u, nil
}

func (r *rdb) UserUpdate(u *users.User) (*users.User, error) {
	u, err := userUpdate(r.db, model.NewUser(u))
	if err != nil {
		return nil, fmt.Errorf("rdb-repository.rdb.UserUpdate: %w", err)
	}
	return u, nil
}

func (r *rdb) UserDelete(userid users.ID) error {
	err := userDelete(r.db, model.NewUserFromID(userid))
	if err != nil {
		return fmt.Errorf("rdb-repository.rdb.UserDelete: %w", err)
	}
	return nil
}
