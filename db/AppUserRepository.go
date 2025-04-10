package db

import (
	"context"
	"forumproject/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type UserRepository interface {
	CreateUser(user *models.AppUser) (*models.AppUser, error)
	GetUser(id string) (*models.AppUser, error)
}

type PostgresUserRepository struct {
	db *Database
}

func NewDbUserRepository(db *Database) (*PostgresUserRepository, error) {
	return &PostgresUserRepository{db: db}, nil
}

func (repo PostgresUserRepository) CreateUser(user *models.AppUser) (*models.AppUser, error) {
	sql := `
	INSERT INTO forum.app_users (uuid, username, password)
	VALUES ($1, $2, $3)
	`
	_, err := repo.db.Pool.Exec(context.Background(), sql, user.Uuid, user.Username, user.GetPassword())
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (repo PostgresUserRepository) UpdateUser(user *models.AppUser) (*models.AppUser, error) {
	sql := `
	UPDATE forum.app_users SET username = $1, password = $2 WHERE uuid = $3
	`
	_, err := repo.db.Pool.Exec(context.Background(), sql, user.Username, user.GetPassword(), user.Uuid)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (repo PostgresUserRepository) GetUserByUsername(username string) (*models.AppUser, error) {
	sql := `
	SELECT uuid, username, password
	FROM forum.app_users
	WHERE username = $1
	`
	rows, err := repo.db.Pool.Query(context.Background(), sql, username)
	if err != nil {
		return nil, err
	}

	var user models.AppUser

	if err := pgxscan.ScanOne(&user, rows); err != nil {
		return nil, err
	}

	return &user, nil
}
