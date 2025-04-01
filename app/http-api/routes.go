package http_api

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/http-api/controllers"
)

func (a *AppContext) MapRoute() {
	// default Routes
	a.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "fine"})
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

	api := a.app.Group("/api", apiMiddleware)

	// Stateless Controller
	controllers.ConfigUserController(a.Config).MapRoute(api.Group("/user"))

	// Stateful Controller
	controllers.NewDemoController(a.Config, a.WS).MapRoute(api.Group("/demo"))
}

func apiMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
