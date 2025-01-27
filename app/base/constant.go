package base

import "github.com/gofiber/fiber/v2"

type ErrorCode int
type FieldInvalidCode int

const (
	NotFound      ErrorCode = 1   // Error Data Notfound
	BadRequest    ErrorCode = 2   // Error Conflict
	Invalidate    ErrorCode = 3   // Error Basic Validate
	Integration   ErrorCode = 4   // Error When Integrate External Service
	UnHandleError ErrorCode = 500 // Internal Server Error, UnHandle Error
)

const (
	ValueCannotBeNull  FieldInvalidCode = 1000 // Error Required
	ValueIsRequired    FieldInvalidCode = 1001 // Error Required
	ValueNotInScope    FieldInvalidCode = 1002 // Error Required
	ValueInvalidFormat FieldInvalidCode = 1003 // Error Required
)

func (e ErrorCode) GetHttpCode() int {
	switch e {
	case NotFound:
		return fiber.StatusNotFound
	case BadRequest:
	case Invalidate:
		return fiber.StatusBadRequest
	case Integration:
	case UnHandleError:
	default:
		return fiber.StatusInternalServerError
	}
	return fiber.StatusInternalServerError
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
	case UnHandleError:
		return "Something went wrong"
	default:
		return "Something went wrong"
	}
}

func (f FieldInvalidCode) GetErrorMsg() string {
	switch f {
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
