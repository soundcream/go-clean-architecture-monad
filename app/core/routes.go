package core

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/core/controllers"
)

func (a *AppContext) MapRoute() {
	// default Routes
	a.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	a.app.Get("/master", func(c *fiber.Ctx) error {
		var data = []string{"a", "b"}
		return c.JSON(fiber.Map{
			"status":  "0",
			"message": "",
			"data":    data,
			"values":  []string{"x", "y"},
		})
	})

	api := a.app.Group("/api", middleware)

	// Stateless Controller
	controllers.ConfigUserController(a.Config).MapRoute(api.Group("/user"))

	// Stateful Controller
	controllers.NewDemoController(a.Config).MapRoute(api.Group("/demo"))
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}
