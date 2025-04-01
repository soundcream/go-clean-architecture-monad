package validators

import (
	"github.com/go-playground/validator/v10"
	"n4a3/clean-architecture/app/core/global"
)

type (
	XValidator struct {
		Validator *validator.Validate
	}
)

func (v XValidator) Validate(data interface{}) []global.InvalidateField {
	var validationErrors []global.InvalidateField
	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elem := global.InvalidateField{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
			}
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
