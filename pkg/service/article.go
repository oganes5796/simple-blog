package service

import (
	"github.com/oganes5796/simple-blog/models"
	"github.com/oganes5796/simple-blog/pkg/repository"
)

type ArticleService struct {
	ArticleRepo repository.Article
}

func NewArticleSerice(articleRepo repository.Article) *ArticleService {
	return &ArticleService{ArticleRepo: articleRepo}
}

func (a *ArticleService) CreateArticle(authorID int, title, content string) (models.Article, error) {
	arrticle := models.Article{AuthorID: authorID, Title: title, Content: content}
	err := a.ArticleRepo.CreateArticle(arrticle)

	return arrticle, err
}

func (a *ArticleService) GetArticleByID(id int) (*models.Article, error) {
	return a.ArticleRepo.GetArticle(id)
}
func (a *ArticleService) GetAllArticle() ([]models.Article, error) {
	return a.ArticleRepo.GetArticles()
}

func (a *ArticleService) UpdateArticle(article models.Article) error {
	return a.ArticleRepo.UpdateArticle(article)
}

func (a *ArticleService) DeleteArticle(id, authorID int) error {
	return a.ArticleRepo.DeleteArticle(id, authorID)
}
