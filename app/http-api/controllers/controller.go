package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/core"
	stringutil "n4a3/clean-architecture/app/core/util/string"
	"n4a3/clean-architecture/app/integrates/dto"
)

type Controller interface {
	MapRoute(route fiber.Router)
}

func MapBody[T any](c *fiber.Ctx) core.Either[T, core.ErrContext] {
	result := new(T)
	if err := c.BodyParser(result); err != nil {
		return core.LeftEither[T, core.ErrContext](core.NewErrorCodeWithMsg(core.BadRequest, "Body empty"))
	}
	return core.RightEither[T, core.ErrContext](*result)
}

// MapCommandByRouteParamsId
func MapCommandByRouteParamsId[T any](c *fiber.Ctx) core.Either[dto.CommandDto[T], core.ErrContext] {
	result := new(dto.CommandDto[T])
	model := new(T)
	if err := c.BodyParser(model); err != nil {
		return core.LeftEither[dto.CommandDto[T], core.ErrContext](core.NewErrorCodeWithMsg(core.BadRequest, "Body empty"))
	}
	if model == nil {
		return core.LeftEither[dto.CommandDto[T], core.ErrContext](core.NewErrorCodeWithMsg(core.BadRequest, "Model is null"))
	}
	result.Model = model
	entityId := stringutil.ToIntEither(GetRouteParams(c, "id"))
	if entityId.IsRight() {
		result.Id = *entityId.Right
	} else {
		return core.LeftEither[dto.CommandDto[T], core.ErrContext](*entityId.Left)
	}
	return core.RightEither[dto.CommandDto[T], core.ErrContext](*result)
}

func GetRouteParamsById(c *fiber.Ctx) core.Either[int, core.ErrContext] {
	entityId := stringutil.ToIntEither(GetRouteParams(c, "id"))
	if entityId.IsRight() {
		return core.RightEither[int, core.ErrContext](*entityId.Right)
	}
	return core.LeftEither[int, core.ErrContext](*entityId.Left)
}

func GetRouteParams(c *fiber.Ctx, routeParamName string) string {
	entityId := c.Params(routeParamName)
	return entityId
}

func Response[T any](c *fiber.Ctx, result core.Either[T, core.ErrContext]) error {
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}

func ErrorHandleResult[T any](c *fiber.Ctx, result core.Either[T, core.ErrContext]) error {
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}

func ErrorResult(c *fiber.Ctx, error *core.ErrContext) error {
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
