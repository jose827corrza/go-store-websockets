package dtos

type ProductRequest struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
	Stock       int     `validate:"required"`
	Image       string  `validate:"required"`
	BrandID     string  `validate:"required"`
	CategoryID  string  `validate:"required"`
}
