package validators

import (
	"github.com/go-playground/validator/v10"
	"n4a3/clean-architecture/app/core"
)

func RegisterIsTeenValidator(validate *validator.Validate) error {
	err := validate.RegisterValidation(string(core.ValidateTeen), func(fl validator.FieldLevel) bool {
		// User.Age needs to fit our needs, 12-18 years old.
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
	if err != nil {
		return err
	}
	return nil
}
