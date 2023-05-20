package validators

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jose827corrza/go-store-websockets/dtos"
)

func ValidateTask(structure *dtos.TaskRequest, w http.ResponseWriter, r *http.Request) error {
	Validate = validator.New()
	// var test dtos.BrandRequest
	err := json.NewDecoder(r.Body).Decode(&structure)
	log.Print(structure)
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

func ValidateUpdateTask(structure *dtos.TaskUpdate, w http.ResponseWriter, r *http.Request) error {
	Validate = validator.New()
	// var test dtos.BrandRequest
	err := json.NewDecoder(r.Body).Decode(&structure)
	log.Print(structure)
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

func ValidateSubTask(structure *dtos.SubTask, w http.ResponseWriter, r *http.Request) error {
	Validate = validator.New()
	// var test dtos.BrandRequest
	err := json.NewDecoder(r.Body).Decode(&structure)
	log.Print(structure)
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
