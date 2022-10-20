package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/rudikurniawan99/go-api-4/src/model"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func UserValidator(req *model.UserRequest) []error {
	validate := validator.New()
	trans := initTranslation(validate)

	if err := validate.Struct(req); err != nil {
		// errs := translateError(err, trans)
		// return errs

		return translateError(err, trans)
	}
	return nil
}

func initTranslation(validate *validator.Validate) ut.Translator {
	english := en.New()

	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

func translateError(err error, trans ut.Translator) (errs []error) {

	validationErrors := err.(validator.ValidationErrors)
	for _, e := range validationErrors {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}

	return errs
}
