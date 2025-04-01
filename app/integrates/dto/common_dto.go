package dto

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
	"time"
)

type Response[T any] struct {
	Data             *T
	Success          bool
	ResponseDateTime time.Time
	Error            *global.ErrorHandlerResp
}

func SuccessResponse[T any](data *T) Response[T] {
	return Response[T]{
		Data:             data,
		Success:          true,
		ResponseDateTime: time.Now(),
	}
}

func ErrorResponse(error global.ErrorHandlerResp) Response[string] {
	return Response[string]{
		Success:          false,
		ResponseDateTime: time.Now(),
		Error:            &error,
	}
}

func ErrorContextResponse(error base.ErrContext) Response[string] {
	return Response[string]{
		Success:          false,
		ResponseDateTime: time.Now(),
		Error: &global.ErrorHandlerResp{
			Code:    int(error.Code),
			Message: error.Code.GetDefaultErrorMsg(),
			Ext:     error.Extensions,
		},
	}
}

func ErrorUnHandlerResponse() Response[string] {
	return Response[string]{
		Success:          false,
		ResponseDateTime: time.Now(),
		Error: &global.ErrorHandlerResp{
			Code:    int(base.UnHandleError),
			Message: base.UnHandleError.GetDefaultErrorMsg(),
		},
	}
}

type CommandDto[T any] struct {
	Id    int `json:"id" example:"1"`
	Model *T  `json:"model"`
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
