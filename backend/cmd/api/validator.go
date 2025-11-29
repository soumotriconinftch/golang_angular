package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type CreateUserPayload struct {
	Name string `json:"name" validate:"required,min=3"`
	Age  int    `json:"age" validate:"required,min=1"`
}

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

func Validator(w http.ResponseWriter, r *http.Request) {
	sampleReq := &CreateUserPayload{}
	name := chi.URLParam(r, "name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "name is required",
		})
		return
	}
	// Parse and decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(sampleReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid JSON format",
		})
		return
	}

	// Validate the request
	if err := Validate.Struct(sampleReq); err != nil {
		validationErrors := make(map[string][]string)
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			validationErrors[fieldName] = append(validationErrors[fieldName], err.Tag())
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Validation error",
			"body":    ValidationErrors{Errors: validationErrors},
		})
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Request received successfully",
		"body":    sampleReq,
	})
}
