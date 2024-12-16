package models

import "time"

type Article struct {
	ID        int       `json:"id" db:"id"`
	AuthorID  int       `json:"author_id" binding:"author_id"`
	Title     string    `json:"title" binding:"title"`
	Content   string    `json:"content" binding:"content"`
	CreatedAt time.Time `json:"created_at" binding:"created_at"`
}
