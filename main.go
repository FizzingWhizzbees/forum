package main

import (
	"context"
	"fmt"
	"forumproject/config"
	"forumproject/db"
	"forumproject/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbConfig := config.DbConfig{DatabaseUrl: "postgres://golang:securePass1@localhost:6500/forum"}
	database := db.NewDatabase(dbConfig)

	pool := database.Pool

	ctx := context.Background()

	initDb(ctx, pool, "DROP TABLE IF EXISTS forum.app_users")
	initDb(ctx, pool, "CREATE TABLE forum.app_users (uuid varchar(255) PRIMARY KEY, username varchar(255), password varchar(255))")

	stat := pool.Stat()
	println("stat:", stat)

	userRepo, _ := db.NewDbUserRepository(database)
	user := models.NewAppUser("piotr", "qwertyuiop123!")
	userRepo.CreateUser(user)
	retrieved, _ := userRepo.GetUserByUsername("piotr")
	fmt.Printf("%#v\n", retrieved)
	println("retrieved:", retrieved)

	retrieved.Username = "piotrek1"
	retrieved.SetPassword("newpassword1")
	userRepo.UpdateUser(retrieved)

	updated, _ := userRepo.GetUserByUsername("piotrek1")
	passcheck, _ := updated.CheckPassword("newpassword1")
	if passcheck {
		println("Password update test passed!")
	}

}

func initDb(ctx context.Context, pool *pgxpool.Pool, sql string) {
	_, err := pool.Exec(ctx, sql)
	if err != nil {
		panic(err)
	}
}
