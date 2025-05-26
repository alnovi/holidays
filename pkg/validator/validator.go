package validator

import (
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type EchoValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewEchoValidator() *EchoValidator {
	validate := validator.New()
	translator := newTranslator(validate)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0] //nolint:mnd
		if name == "" {
			name = field.Tag.Get("form")
		}

		if name == "-" {
			return ""
		}

		return name
	})

	return &EchoValidator{
		validator:  validate,
		translator: translator,
	}
}

func (e *EchoValidator) Validate(i interface{}) error {
	if err := e.validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		return e.convertErrors(err)
	}

	return nil
}

func (e *EchoValidator) AddRule(rule Rule) error {
	if err := e.validator.RegisterValidation(rule.Tag(), rule.Validate, rule.CallIfNull()); err != nil {
		return err
	}

	return e.validator.RegisterTranslation(rule.Tag(), e.translator, func(ut ut.Translator) error {
		return ut.Add(rule.Tag(), rule.ErrMsg(), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(rule.Tag(), fe.Field())
		return t
	})
}

func (e *EchoValidator) convertErrors(validErrors error) *ValidateError {
	err := NewValidateError()

	for _, fe := range validErrors.(validator.ValidationErrors) {
		namespace := strings.SplitN(fe.Namespace(), ".", 2) //nolint:mnd
		if len(namespace) == 1 {
			err.Fields[namespace[0]] = fe.Translate(e.translator)
		} else {
			err.Fields[namespace[1]] = fe.Translate(e.translator)
		}
	}

	return err
}
