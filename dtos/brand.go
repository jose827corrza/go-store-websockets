package dtos

type BrandRequest struct {
	Name  string `validate:"required"`
	Image string `validate:"required"`
}
