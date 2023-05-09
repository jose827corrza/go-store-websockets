package dtos

type CategoryRequest struct {
	Name  string `validate:"required"`
	Image string `validate:"required"`
}
