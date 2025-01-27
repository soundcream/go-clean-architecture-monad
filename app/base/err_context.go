package base

import "n4a3/clean-architecture/app/base/util"

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

func NewBadError() ErrContext {
	errorCode := BadRequest
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      errorCode.GetDefaultErrorMsg(),
	}
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

func NewInvalidateError(field string, code FieldInvalidCode) *ErrContext {
	return NewInvalidateErrorWithMsg(field, code, code.GetErrorMsg())
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

func NewError(errorCode ErrorCode, err error) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      errorCode.GetDefaultErrorMsg(),
		Cause:    getCause(err),
		error:    err}
}

func NewErrorWithMsg(errorCode ErrorCode, msg string, err error) ErrContext {
	return ErrContext{
		Code:     errorCode,
		HttpCode: errorCode.GetHttpCode(),
		Msg:      msg,
		Cause:    getCause(err),
		error:    err}
}

func getCause(err error) *string {
	if err == nil {
		return nil
	}
	return util.ToPtr(err.Error())
}
