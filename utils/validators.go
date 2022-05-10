package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}