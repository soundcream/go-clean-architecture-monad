package core

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorCode int
type FieldInvalidCode int
type ValidateType string

const (
	Environment = "env"
)

// ErrorCode Error Code
const (
	NotFound      ErrorCode = 404 // Error Data Notfound
	BadRequest    ErrorCode = 400 // Error Invalid Model or Model is null
	Invalidate    ErrorCode = 410 // Error Invalidate data (Handle)
	Conflict      ErrorCode = 411 // Error Conflict data conflict (Handle)
	Integration   ErrorCode = 412 // Error When Integrate External Service
	Invalid       ErrorCode = 413 // Error Not complete
	Unauthorized  ErrorCode = 401 // Error Unauthorized
	UnHandleError ErrorCode = 500 // Internal Server Error, UnHandle Error
)

// FieldInvalidCode Error Code
const (
	ValueInvalidate    FieldInvalidCode = 1000 // Error Required
	ValueCannotBeNull  FieldInvalidCode = 1001 // Error Required
	ValueIsRequired    FieldInvalidCode = 1002 // Error Required
	ValueNotInScope    FieldInvalidCode = 1003 // Error Required
	ValueInvalidFormat FieldInvalidCode = 1004 // Error Required
)

// Validation Tag
const (
	ValidateTeen          ValidateType = "teen-person"
	ValidateEmailUsername ValidateType = "username-email"
)

func (e ErrorCode) GetHttpCode() int {
	switch e {
	case NotFound:
		return fiber.StatusNotFound
	case BadRequest:
		return fiber.StatusBadRequest
	case Invalidate:
		return fiber.StatusBadRequest
	case Integration:
		return fiber.StatusInternalServerError
	case Unauthorized:
		return fiber.StatusUnauthorized
	case UnHandleError:
		return fiber.StatusInternalServerError
	default:
		return fiber.StatusInternalServerError
	}
}

func (e ErrorCode) GetDefaultErrorMsg() string {
	switch e {
	case NotFound:
		return "Not resource matched"
	case BadRequest:
		return "Invalid request"
	case Invalidate:
		return "Invalid request"
	case Integration:
		return "Something went wrong"
	case Conflict:
		return "Invalid data request"
	case Invalid:
		return "Not complete"
	case Unauthorized:
		return "Unauthorized"
	case UnHandleError:
		return "Something went wrong"
	default:
		return "Something went wrong"
	}
}

func (f FieldInvalidCode) GetErrorMsg() string {
	switch f {
	case ValueInvalidate:
		return "Value is invalid"
	case ValueIsRequired:
		return "Value is required"
	case ValueNotInScope:
		return "Value not in scope"
	case ValueInvalidFormat:
		return "Invalid format"
	case ValueCannotBeNull:
		return "Value cannot be null"
	default:
		return "Value is required"
	}
}

func (v ValidateType) getMessage() string {
	switch v {
	case ValidateTeen:
		return ""
	case ValidateEmailUsername:
		return ""
	default:
		return "-"
	}
}
