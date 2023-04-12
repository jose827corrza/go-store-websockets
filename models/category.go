package models

import "time"

type Category struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Products  []Product `json:"products"`
}
