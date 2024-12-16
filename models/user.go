package models

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Username  string    `json:"username" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
