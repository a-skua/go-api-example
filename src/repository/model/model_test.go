package model

import (
	"api.example.com/env"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"sync"
	"testing"
	"time"
)

var tableLock sync.Mutex

type queryResult struct {
	lastID int64
	rows   int64
	err    error
	// flag
	lastInsertID, rowsAffected bool
}

func (result *queryResult) LastInsertId() (int64, error) {
	if result.lastInsertID {
		return result.lastID, result.err
	}
	return 0, errors.New("test invalid LastInsertId")
}

func (result *queryResult) RowsAffected() (int64, error) {
	if result.rowsAffected {
		return result.rows, result.err
	}
	return 0, errors.New("test invalid RowsAffected")
}

type testdb struct {
	result sql.Result
	err    error
	// flag
	execContext bool
}

func (db *testdb) ExecContext(context.Context, string, ...interface{}) (sql.Result, error) {
	if db.execContext {
		return db.result, db.err
	}
	return nil, errors.New("test invalid ExecContext")
}

func (db *testdb) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	panic("test invalid QueryRowContext")
}

func newDB() *sql.DB {

	addr := env.Get("TEST_DB_ADDR")
	name := env.Get("TEST_DB_NAME")
	user := env.Get("TEST_DB_USER")
	password := env.GetSecure("TEST_DB_PASSWORD")

	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true",
		user.Value(),
		password.Value(),
		addr.Value(),
		name.Value(),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func testDiffTime(t *testing.T, want, got time.Time) {
	t.Helper()

	if math.Abs(float64(want.Sub(got))) > float64(time.Second) {
		// 1秒以上ずれる場合はエラーとする
		t.Fatalf("updated_at want=%v, got=%v.", want, got)
	}
}

func TestCurrentTime(t *testing.T) {
	want := time.Now().Round(time.Second)
	got := currentTime()

	testDiffTime(t, want, got)
}
