package helper

import "github.com/go-playground/validator/v10"

func ValidateObj(object interface{}) error {
	validate := validator.New()
	return validate.Struct(object)
}
