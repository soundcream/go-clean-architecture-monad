package dto

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
)

type Response[T any] struct {
	Data    *T
	Success bool
	Error   global.ErrorHandlerResp
}

func SuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Data:    &data,
		Success: true,
	}
}

func ErrorResponse(error global.ErrorHandlerResp) Response[string] {
	return Response[string]{
		Success: false,
		Error:   error,
	}
}

func ErrorContextResponse(error base.ErrContext) Response[string] {
	return Response[string]{
		Success: false,
		Error: global.ErrorHandlerResp{
			Code:    int(error.Code),
			Message: error.Code.GetDefaultErrorMsg(),
			Ext:     error.Extensions,
		},
	}
}

func ErrorUnHandlerResponse() Response[string] {
	return Response[string]{
		Success: false,
		Error: global.ErrorHandlerResp{
			Code:    int(base.UnHandleError),
			Message: base.UnHandleError.GetDefaultErrorMsg(),
		},
	}
}

type PagingDto struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PaginationDto[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
