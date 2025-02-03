package global

import (
	"github.com/go-playground/validator/v10"
)

func NewKeyValue[T, S any](key T, value S) KeyValue[T, S] {
	return KeyValue[T, S]{Key: key, Val: value}
}

type (
	Config struct {
		App          AppConfig
		DbConfig     DbConfig
		ReadDbConfig DbConfig
		Public       []string
	}
	AppConfig struct {
		AppName  string
		Domain   string
		HttpPort int
	}
	DbConfig struct {
		Host     string
		Port     int
		DbName   string
		Username string
		Password string
	}
	User struct {
		Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
		Age  int    `validate:"required,teen-person"`  // Required field, and client needs to implement our 'teen-person' tag format which we'll see later
	}
	InvalidateField struct {
		Error       bool
		FailedField string
		Tag         string
		Msg         string
		Value       interface{}
	}
	XValidator struct {
		Validator *validator.Validate
	}
	ErrorHandlerResp struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Ext     interface{} `json:"ext"`
	}
	KeyValue[T, S any] struct {
		Key T
		Val S
	}
)

func (v XValidator) Validate(data interface{}) []InvalidateField {
	validationErrors := []InvalidateField{}
	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elem := InvalidateField{
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

//func (v ErrorContextValidator) ValidateContext(data interface{}) *base.ErrContext {
//	validationErrors := []InvalidateField{}
//	var validate = validator.New()
//	errs := validate.Struct(data)
//	if errs != nil {
//		for _, err := range errs.(validator.ValidationErrors) {
//			// In this case data object is actually holding the User struct
//			var elem InvalidateField
//			elem.FailedField = err.Field() // Export struct field name
//			elem.Tag = err.Tag()           // Export struct tag
//			elem.Value = err.Value()       // Export field value
//			elem.Error = true
//			validationErrors = append(validationErrors, elem)
//		}
//	}
//
//	return base.NewErrContextFromInvalidateField(validationErrors)
//}
