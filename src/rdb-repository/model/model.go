package model

import (
	"context"
	"database/sql"
)

// SQL 抽象化
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
