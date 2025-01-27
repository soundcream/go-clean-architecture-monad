package base

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Data T
	Code int
	Msg  string
	Ext  map[string]string
}

func NewResponse[T any](data T) Response[T] {
	return Response[T]{
		Data: data,
	}
}

func (r Response[T]) OK(c *fiber.Ctx) {

}
