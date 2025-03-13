package base

import (
	"n4a3/clean-architecture/app/base/collection"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/base/util"
)

type ErrContext struct {
	Code       ErrorCode
	HttpCode   int
	Msg        string
	Cause      *string
	Extensions *[]ErrExt
	error
}

type ErrExt struct {
	Code  int
	Field string
	Msg   string
}

func (e *ErrContext) AppendExt(err *ErrContext) *ErrContext {
	if e.Extensions != nil && err != nil {
		ex := append(*e.Extensions, *err.Extensions...)
		e.Extensions = &ex
	} else if err != nil {
		e.Extensions = err.Extensions
	}
	return e
}

func NewIfError(err error) *ErrContext {
	if err == nil {
		return nil
	}
	return &ErrContext{
		Code:     UnHandleError,
		HttpCode: UnHandleError.GetHttpCode(),
		Msg:      UnHandleError.GetDefaultErrorMsg(),
		Cause:    getCause(err),
		error:    err,
	}
}

func NewErrorCode(errorCode ErrorCode) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      errorCode.GetDefaultErrorMsg(),
	}
}

func NewErrorCodeWithMsg(errorCode ErrorCode, cause string) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      errorCode.GetDefaultErrorMsg(),
		Cause:    &cause,
	}
}

func NewErrorWithCode(errorCode ErrorCode, err error) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      errorCode.GetDefaultErrorMsg(),
		Cause:    getCause(err),
		error:    err,
	}
}

func NewErrorWithMsg(errorCode ErrorCode, msg string, err error) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      msg,
		Cause:    getCause(err),
		error:    err}
}

func NewInvalidateExtError(ext []ErrExt) ErrContext {
	errorCode := Invalidate
	return ErrContext{
		Extensions: &ext,
		Code:       errorCode,
		HttpCode:   errorCode.GetHttpCode(),
		Msg:        errorCode.GetDefaultErrorMsg(),
	}
}

func NewErrContextFromInvalidateField(invalidates []global.InvalidateField) *ErrContext {
	if len(invalidates) == 0 {
		return nil
	}
	result := NewInvalidateExtError(
		collection.NewMapping[global.InvalidateField, ErrExt](invalidates).Map(func(field global.InvalidateField) ErrExt {
			return ErrExt{
				Code:  int(ValueInvalidate),
				Field: field.FailedField,
				Msg:   ValueInvalidate.GetErrorMsg(),
			}
		}))
	return &result
}

func NewInvalidateError(field string, code FieldInvalidCode) *ErrContext {
	return NewInvalidateErrorWithMsg(field, code, code.GetErrorMsg())
}

func NewInvalidateErrorWithMsg(field string, code FieldInvalidCode, msg string) *ErrContext {
	errorCode := Invalidate
	ext := ErrExt{
		Code:  int(code),
		Field: field,
		Msg:   msg,
	}
	return &ErrContext{
		Extensions: &([]ErrExt{ext}),
		Code:       errorCode,
		HttpCode:   errorCode.GetHttpCode(),
		Msg:        errorCode.GetDefaultErrorMsg(),
	}
}

func getCause(err error) *string {
	if err == nil {
		return nil
	}
	return util.ToPtr(err.Error())
}
