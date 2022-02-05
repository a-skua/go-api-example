package repository

import (
	"api.example.com/env"
	"database/sql"
	"fmt"
	"testing"
)

// テスト用DB
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

func TestUserCreate(t *testing.T) {
	// TODO
}

func TestUserRead(t *testing.T) {
	// TODO
}

func TestUserUpdate(t *testing.T) {
	// TODO
}

func TestUserDelete(t *testing.T) {
	// TODO
}
