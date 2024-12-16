package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oganes5796/simple-blog/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password, role) VALUES ($1, $2, $3) RETURNING id", usersTable)
	if err := a.db.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.Id); err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}
