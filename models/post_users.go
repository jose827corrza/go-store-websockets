package models

import "gorm.io/gorm"

type UserPost struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	ID       string
	Posts    []Post // 1:N
}
