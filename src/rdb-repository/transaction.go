package repository

import (
	"api.example.com/pkg/repository"
	"database/sql"
	"errors"
)

type Tx struct {
	sql *sql.Tx
	err error
}

func newTx(sql *sql.Tx, err error) repository.Tx {
	return &Tx{
		sql: sql,
		err: err,
	}
}

func castTx(tran repository.Tx) (*Tx, error) {
	tx, ok := tran.(*Tx)
	if !ok {
		return nil, errors.New("failed cast to *Tx")
	}
	return tx, nil
}

func (tx *Tx) Rollback() error {
	if tx == nil {
		return errors.New("*Tx is nil")
	}
	if tx.sql == nil {
		return errors.New("*Tx.sql is nil")
	}
	return tx.sql.Rollback()
}

func (tx *Tx) Commit() error {
	if tx == nil {
		return errors.New("*Tx is nil")
	}
	if tx.sql == nil {
		return errors.New("*Tx.slq is nil")
	}
	return tx.sql.Commit()
}

func (tx *Tx) Error() error {
	if tx == nil {
		return errors.New("*Tx is nil")
	}
	return tx.err
}

func (tx *Tx) newState(err error) *Tx {
	var sqlTx *sql.Tx
	if tx != nil {
		sqlTx = tx.sql
	}
	return &Tx{sqlTx, err}
}
