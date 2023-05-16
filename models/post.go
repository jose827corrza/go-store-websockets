package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string
	Description string
	ID          string
	UserID      string // 1:N
}
