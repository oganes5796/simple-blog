package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	ReaderID  int       `json:"reader_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
