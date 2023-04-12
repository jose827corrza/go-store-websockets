package models

import "time"

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Customer  Customer  `json:"customer,omitempty"`
}
type RoleTest string

const (
	ADMIN    RoleTest = "administrator"
	CUSTOMER RoleTest = "customer"
)

type Role struct {
	Administrator RoleTest
}
