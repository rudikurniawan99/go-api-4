package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/rudikurniawan99/go-api-4/src/model"
)

func UserValidator(req *model.UserRequest) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}
