package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
)

type DemoController struct {
	Facade facades.DemoFacade
}

func NewDemoController() *DemoController {
	return &DemoController{
		Facade: facades.NewDemoFacade(),
	}
}

func (con *DemoController) MapRoute(route fiber.Router) {
	route.Get("/ex", func(c *fiber.Ctx) error {
		return con.TestValidate(c)
	})
}

// TestValidate @Summary Example of chain Validate
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/ex [get]
func (con *DemoController) TestValidate(c *fiber.Ctx) error {
	result := con.Facade.Validate(new(entity.User))
	if result.IsRight() {
		//return core.OkResult(c, result.Right)
		return c.Status(fiber.StatusOK).JSON(result)
	}
	//return core.ErrorResult(c, result.Left)
	return c.Status(fiber.StatusInternalServerError).JSON(result)
}
