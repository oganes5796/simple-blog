package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/oganes5796/simple-blog/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Article interface {
	CreateArticle(article models.Article) error
	GetArticle(id int) (*models.Article, error)
	GetArticles() ([]models.Article, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id, authorID int) error
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Article:       NewArticleRepository(db),
	}
}
