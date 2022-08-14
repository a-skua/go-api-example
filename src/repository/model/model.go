package model

import (
	"context"
	"database/sql"
	"time"
)

// SQL 抽象化
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type dateTime = time.Time

func currentTime() dateTime {
	return time.Now().Round(time.Second)
}
