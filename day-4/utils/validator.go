package utils

import (
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// GetValidator returns a custom validator
func GetValidator() *CustomValidator {
	cv := &CustomValidator{validator: validator.New()}
	return cv
}
