package dtos

type PostRequest struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}
