package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
	Posts    []Post
	Name     string `json:"name"`
	// CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Customer  Customer  `json:"customer,omitempty"`
}
type RoleTest string

const (
	ADMIN    RoleTest = "administrator"
	CUSTOMER RoleTest = "customer"
)

type Role struct {
	Administrator RoleTest
}
