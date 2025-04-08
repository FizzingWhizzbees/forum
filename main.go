package main

import (
	"context"
	"forumproject/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljamescleary/gomin/pkg/common/config"
)

func main() {
	pool, err := db.NewPgConnectionPool(config.Config{DbUrl: "postgres://golang:securePass1@localhost:6500/forum"})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	initDb(ctx, pool, "DROP TABLE IF EXISTS forum.app_users")
	initDb(ctx, pool, "CREATE TABLE forum.app_users (uuid varchar(255) PRIMARY KEY, username varchar(255), password varchar(255))")

	stat := pool.Stat()
	println("stat:", stat)
}

func initDb(ctx context.Context, pool *pgxpool.Pool, sql string) {
	_, err := pool.Exec(ctx, sql)
	if err != nil {
		panic(err)
	}
}
