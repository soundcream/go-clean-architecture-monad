package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/base"
	stringutil "n4a3/clean-architecture/app/base/util/string"
	"n4a3/clean-architecture/app/integrates/dto"
)

type Controller interface {
	MapRoute(route fiber.Router)
}

func MapBody[T any](c *fiber.Ctx) base.Either[T, base.ErrContext] {
	result := new(T)
	if err := c.BodyParser(result); err != nil {
		return base.LeftEither[T, base.ErrContext](base.NewErrorCodeWithMsg(base.BadRequest, "Body empty"))
	}
	return base.RightEither[T, base.ErrContext](*result)
}

// MapCommandByRouteParamsId
func MapCommandByRouteParamsId[T any](c *fiber.Ctx) base.Either[dto.CommandDto[T], base.ErrContext] {
	result := new(dto.CommandDto[T])
	model := new(T)
	if err := c.BodyParser(model); err != nil {
		return base.LeftEither[dto.CommandDto[T], base.ErrContext](base.NewErrorCodeWithMsg(base.BadRequest, "Body empty"))
	}
	if model == nil {
		return base.LeftEither[dto.CommandDto[T], base.ErrContext](base.NewErrorCodeWithMsg(base.BadRequest, "Model is null"))
	}
	result.Model = model
	entityId := stringutil.ToIntEither(GetRouteParams(c, "id"))
	if entityId.IsRight() {
		result.Id = *entityId.Right
	} else {
		return base.LeftEither[dto.CommandDto[T], base.ErrContext](*entityId.Left)
	}
	return base.RightEither[dto.CommandDto[T], base.ErrContext](*result)
}

func GetRouteParamsById(c *fiber.Ctx) base.Either[int, base.ErrContext] {
	entityId := stringutil.ToIntEither(GetRouteParams(c, "id"))
	if entityId.IsRight() {
		return base.RightEither[int, base.ErrContext](*entityId.Right)
	}
	return base.LeftEither[int, base.ErrContext](*entityId.Left)
}

func GetRouteParams(c *fiber.Ctx, routeParamName string) string {
	entityId := c.Params(routeParamName)
	return entityId
}

func Response[T any](c *fiber.Ctx, result base.Either[T, base.ErrContext]) error {
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
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

func OkResult[T any](c *fiber.Ctx, data *T) error {
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
