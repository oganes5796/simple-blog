package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/oganes5796/simple-blog/models"
)

type ArticlesRepository struct {
	DB *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticlesRepository {
	return &ArticlesRepository{DB: db}
}

func (r *ArticlesRepository) CreateArticle(article models.Article) error {
	query := "INSERT INTO articles (author_id, title, content) VALUES ($1, $2, $3) RETURNING id, created_at"
	if err := r.DB.QueryRow(query, article.AuthorID, article.Title, article.Content).Scan(&article.ID, &article.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *ArticlesRepository) GetArticle(id int) (*models.Article, error) {
	query := "SELECT id, author_id, title, content, created_at FROM posts WHERE id = $1"
	post := &models.Article{}
	err := r.DB.Get(post, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return post, nil
}
func (r *ArticlesRepository) GetArticles() ([]models.Article, error) {
	query := "SELECT id, author_id, title, content, created_at FROM posts"
	articles := []models.Article{}
	err := r.DB.Select(&articles, query)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
func (r *ArticlesRepository) UpdateArticle(article models.Article) error {
	query := "UPDATE Articles SET title = $1, content = $2 WHERE id = $3 AND author_id = $4"
	result, err := r.DB.Exec(query, article.Title, article.Content, article.ID, article.AuthorID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
func (r *ArticlesRepository) DeleteArticle(id, authorID int) error {
	query := "DELETE FROM posts WHERE id = $1 AND author_id = $2"
	result, err := r.DB.Exec(query, id, authorID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
