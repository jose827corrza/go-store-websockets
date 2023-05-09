package dtos

// USING VALIDATE
type SignUpUserRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type SignUpUserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
