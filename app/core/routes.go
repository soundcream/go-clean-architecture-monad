package core

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/core/controllers"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/interfaces/repository"
)

func MapRoute(app *fiber.App) {
	// default Routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	app.Get("/master", func(c *fiber.Ctx) error {
		var data = []string{"a", "b"}
		return c.JSON(fiber.Map{
			"status":  "0",
			"message": "",
			"data":    data,
			"values":  []string{"x", "y"},
		})
	})

	api := app.Group("/api", middleware)

	controllers.MapRoute(api.Group("/user"), getUserController())
}

// TOD use DI
func getUserController() controllers.UserController {
	ur := repository.NewUserRepository()
	uf := facades.NewUserFacade(ur)
	uc := controllers.NewUserController(uf)
	return *uc
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}
