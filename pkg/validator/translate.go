package validator

import (
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ruTranslations "github.com/go-playground/validator/v10/translations/ru"
)

func newTranslator(validate *validator.Validate) ut.Translator {
	ruTranslator := ru.New()
	universalTranslator := ut.New(ruTranslator)
	translator, _ := universalTranslator.GetTranslator("ru")
	_ = ruTranslations.RegisterDefaultTranslations(validate, translator)

	return translator
}
