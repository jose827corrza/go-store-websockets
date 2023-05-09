package dtos

type SignUpCustomerRequest struct {
	Name     string `validate:"required"`
	LastName string
	Phone    string
	Email    string `validate:"required"`
	User     SignUpUserRequest
}
