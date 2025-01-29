package core

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/core/controllers"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/interfaces/repository"
)

func (a *Application) MapRoute() {
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
	MapStateLessUserController(api.Group("/user"))

	// Stateful Controller
	controllers.NewDemoController().MapRoute(api.Group("/demo"))
}

// TOD use DI
func getUserController() controllers.UserController {
	ur := repository.NewUserRepository()
	uf := facades.NewUserFacade(ur)
	uc := controllers.NewUserController(uf)
	return *uc
}

func MapStateLessUserController(route fiber.Router) {
	route.Get("/users", func(c *fiber.Ctx) error {
		con := getUserController()
		return con.GetUsers(c)
	})
	route.Get("/validate", func(c *fiber.Ctx) error {
		con := getUserController()
		return con.TestValidate(c)
	})
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}
