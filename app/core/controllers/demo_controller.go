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

func (con *DemoController) TestValidate(c *fiber.Ctx) error {
	e := con.Facade.Validate(new(entity.User))
	if e.IsLeft() {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": e.Left})
	}
	return c.JSON(e.Right)
}
