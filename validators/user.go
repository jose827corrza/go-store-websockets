package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jose827corrza/go-store-websockets/dtos"
)

var validate *validator.Validate

func ValidateSignUp(structure *dtos.SignUpUserRequest, w http.ResponseWriter, r *http.Request) error {
	validate = validator.New()

	err := json.NewDecoder(r.Body).Decode(&structure)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = validate.Struct(structure)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// utils.ErrorResponse(400, err.Error(), w)
		return err
	}
	return nil
}
