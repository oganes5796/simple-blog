package service

import (
	"github.com/oganes5796/simple-blog/models"
	"github.com/oganes5796/simple-blog/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateJWT(username, password string) (string, error)
	ParseJWT(tokenString string) (int, error)
}

type Article interface {
	CreateArticle(authorID int, title, content string) (models.Article, error)
	GetArticleByID(id int) (*models.Article, error)
	GetAllArticle() ([]models.Article, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id, authorID int) error
}

type Service struct {
	Authorization
	Article
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Article:       NewArticleSerice(repos.Article),
	}
}
