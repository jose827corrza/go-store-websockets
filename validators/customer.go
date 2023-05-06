package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jose827corrza/go-store-websockets/dtos"
)

func ValidateCustomerSignUp(structure *dtos.SignUpCustomerRequest, w http.ResponseWriter, r *http.Request) error {
	Validate = validator.New()

	err := json.NewDecoder(r.Body).Decode(&structure)

	if err != nil {
		return err
	}
	err = Validate.Struct(structure)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// utils.ErrorResponse(400, err.Error(), w)
		return err
	}
	return nil
}
