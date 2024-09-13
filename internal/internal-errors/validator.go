package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]
	fieldName := strings.ToLower(validationError.StructField())

	switch validationError.Tag() {
	case "required":
		return errors.New(fieldName + " is required")
	case "min":
		return errors.New(fieldName + " is less than the minimum: " + validationError.Param())
	case "max":
		return errors.New(fieldName + " is greater than the maximum: " + validationError.Param())
	case "email":
		return errors.New(fieldName + " is not a valid email")
	default:
		return errors.New("unknown validation error on: " + fieldName)
	}
}
