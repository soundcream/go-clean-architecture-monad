package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
)

// UserController is responsible for handling user-related routes.
type UserController struct {
	Facade facades.UserFacade
}

func NewUserController(facade facades.UserFacade) *UserController {
	return &UserController{Facade: facade}
}

func MapRoute(route fiber.Router, controller UserController) {
	route.Get("/users", func(c *fiber.Ctx) error {
		return controller.GetUsers(c)
	})
	route.Get("/validate", func(c *fiber.Ctx) error {
		return controller.TestValidate(c)
	})
}

// GetUsers godoc
// @Summary Get a list of users
// @Description Retrieve all users from the system
// @Tags users
// @Produce json
// @Success 200 {array} entity.User
// @Router /api/user/users [get]
func (u *UserController) GetUsers(c *fiber.Ctx) error {
	u.Facade.TestThen()
	users, err := u.Facade.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (u *UserController) GetUser(c *fiber.Ctx) error {
	users, err := u.Facade.CreateUser("", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (u *UserController) TestValidate(c *fiber.Ctx) error {
	result := u.Facade.TestValidate(&entity.User{})
	return c.Status(fiber.StatusInternalServerError).JSON(result)
}

//
