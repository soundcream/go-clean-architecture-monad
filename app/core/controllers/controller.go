package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/interfaces/dto"
)

type Controller interface {
	MapRoute(route fiber.Router)
}

func ErrorHandleResult[T any](c *fiber.Ctx, result base.Either[T, base.ErrContext]) error {
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}

func ErrorResult(c *fiber.Ctx, error *base.ErrContext) error {
	if error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorUnHandlerResponse())
	}
	return c.Status(error.HttpCode).JSON(dto.ErrorContextResponse(*error))
}

func OkResult(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(data))
}

func Bad(c *fiber.Ctx) {

}

func Notfound(c *fiber.Ctx) {

}

func BadWithI18n(c *fiber.Ctx) {

}

func NotfoundWithI18n(c *fiber.Ctx) {

}

func getError() error {
	var er = errors.New("data not found")
	return er
}
