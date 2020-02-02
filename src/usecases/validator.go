package usecases

import (
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(v interface{}) error
}

func NewValidator() Validator {
	return validator.New()
}
