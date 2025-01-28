package base

import "github.com/gofiber/fiber/v2"

type ErrorCode int
type FieldInvalidCode int

// Request Error
const (
	NotFound      ErrorCode = 100 // Error Data Notfound
	BadRequest    ErrorCode = 101 // Error Invalid Model, null
	Invalidate    ErrorCode = 102 // Error Invalidate
	Conflict      ErrorCode = 103 // Error Conflict
	Integration   ErrorCode = 104 // Error When Integrate External Service
	UnHandleError ErrorCode = 500 // Internal Server Error, UnHandle Error
)

// Field Error
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
		return fiber.StatusBadRequest
	case Invalidate:
		return fiber.StatusBadRequest
	case Integration:
		return fiber.StatusInternalServerError
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
