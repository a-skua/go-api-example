package main

import (
	"api.example.com/env"
	"api.example.com/http-handle"
	"api.example.com/rdb-repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 起動するサーバー本体
var srv http.Server

// データベース
var db *sql.DB

// サーバーの初期化
func init() {
	addr := env.Get("ADDR")
	log.Println("ADDR", addr)

	srv.Addr = addr.Value()
}

// データベースの初期化
func init() {
	addr := env.Get("DB_ADDR")
	name := env.Get("DB_NAME")
	user := env.Get("DB_USER")
	password := env.GetSecure("DB_PASSWORD")
	log.Println("DB_ADDR", addr)
	log.Println("DB_NAME", name)
	log.Println("DB_ADDR", user)
	log.Println("DB_PASSWORD", password)

	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true",
		user.Value(),
		password.Value(),
		addr.Value(),
		name.Value(),
	)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("main SQL Open: %v", err)
	}
}

func main() {
	defer db.Close()
	srv.Handler = handle.New(repository.New(db))

	// 異常終了しないためのおまじない
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	// サーバーの起動
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
