package validator

import "github.com/go-playground/validator/v10"

type Rule interface {
	Tag() string
	ErrMsg() string
	CallIfNull() bool
	Validate(fl validator.FieldLevel) bool
}
