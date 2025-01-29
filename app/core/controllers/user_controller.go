package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
)

// UserController is responsible for handling user-related routes.
type UserController struct {
	Facade facades.UserFacade
	Ctx    *fiber.Ctx
}

func NewUserController(facade facades.UserFacade) *UserController {
	return &UserController{Facade: facade}
}

// GetUsers @Summary Get a list of users
// @Description Retrieve all users from the system
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.User
// @Router /api/user/users [get]
func (u *UserController) GetUsers(c *fiber.Ctx) error {
	u.Facade.TestThen()
	c.Query("take", "10")
	users, err := u.Facade.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (u *UserController) GetUser(c *fiber.Ctx) error {
	user, err := u.Facade.CreateUser("", "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (u *UserController) TestValidate(c *fiber.Ctx) error {
	result := u.Facade.TestValidate(&entity.User{})
	return c.Status(fiber.StatusInternalServerError).JSON(result)
}

func (u *UserController) MapRequestContext(fun func(ctx *fiber.Ctx) error) {

}
