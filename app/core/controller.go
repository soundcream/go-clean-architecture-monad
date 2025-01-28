package core

import "github.com/gofiber/fiber/v2"

type Controller interface {
	MapRoute(route fiber.Router)
}

//func MapController[T any](route fiber.Router, controller UserController) {
//	route.Get("/users", func(c *fiber.Ctx) error {
//		return controller.GetUsers(c)
//	})
//	route.Get("/validate", func(c *fiber.Ctx) error {
//		return controller.TestValidate(c)
//	})
//}
