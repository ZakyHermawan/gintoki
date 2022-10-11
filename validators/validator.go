package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func ValidateCoolTitle(fl validator.FieldLevel) bool {
	println(fl.Field().String())
	println("TRYING TO VALIDATING")
	return false
	return strings.Contains(fl.Field().String(), "Cool")
}
