package repository

import (
	"api.example.com/pkg/repository"
	"database/sql"
	"errors"
)

type transaction struct {
	tx  *sql.Tx
	err error
}

func newTx(tx *sql.Tx, err error) repository.Tx {
	return &transaction{
		tx:  tx,
		err: err,
	}
}

func castTx(tran repository.Tx) (*transaction, error) {
	tx, ok := tran.(*transaction)
	if !ok {
		return nil, errors.New("failed cast tx")
	}
	return tx, nil
}

func (t *transaction) Rollback() error {
	if t == nil {
		return errors.New("*transaction is nil")
	}
	if t.tx == nil {
		return errors.New("*transaction.tx is nil")
	}
	return t.tx.Rollback()
}

func (t *transaction) Commit() error {
	if t == nil {
		return errors.New("*transaction is nil")
	}
	if t.tx == nil {
		return errors.New("*transaction.tx is nil")
	}
	return t.tx.Commit()
}

func (t *transaction) Error() error {
	if t == nil {
		return errors.New("*transaction is nil")
	}
	return t.err
}

func (t *transaction) new(err error) *transaction {
	var tx *sql.Tx
	if t != nil {
		tx = t.tx
	}
	return &transaction{tx, err}
}
