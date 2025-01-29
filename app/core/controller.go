package core

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/base"
)

func OkResult(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse(data))
}

func ErrorResult(c *fiber.Ctx, error *base.ErrContext) error {
	if error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorUnHandlerResponse())
	}
	return c.Status(error.HttpCode).JSON(ErrorContextResponse(*error))
}

func Bad(c *fiber.Ctx) {

}

func Notfound(c *fiber.Ctx) {

}

func OkWithI18n(c *fiber.Ctx) {

}

func BadWithI18n(c *fiber.Ctx) {

}

func NotfoundWithI18n(c *fiber.Ctx) {

}

func getError() error {
	var er = errors.New("data not found")
	return er
}
