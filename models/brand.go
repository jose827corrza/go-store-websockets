package models

import "time"

type Brand struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Products  []Product `json:"products"`
}
