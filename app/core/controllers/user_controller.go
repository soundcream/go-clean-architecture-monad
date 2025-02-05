package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/integrates/db"
	"n4a3/clean-architecture/app/integrates/repository"
)

// UserController is responsible for handling user-related routes.
type UserController struct {
	Facade facades.UserFacade
	Config *global.Config
}

func (u *UserController) MapRoute(route fiber.Router) {
	route.Get("/users", func(c *fiber.Ctx) error {
		return NewUserController(u.Config).GetUsers(c)
	})
	route.Get("/validate", func(c *fiber.Ctx) error {
		return NewUserController(u.Config).TestValidate(c)
	})
}

func ConfigUserController(config *global.Config) *UserController {
	return &UserController{
		Config: config,
	}
}

func NewUserController(config *global.Config) *UserController {
	qUoW := db.NewQueryUnitOfWork(config)
	uow := db.NewUnitOfWork(config)
	//uow := db.NewQueryUnitOfWork(config)
	ur := repository.NewUserRepository(qUoW.Right, uow.Right)
	uf := facades.NewUserFacade(ur)
	return &UserController{
		Config: config,
		Facade: uf,
	}
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
	return OkResult(c, user)
}

func (u *UserController) TestValidate(c *fiber.Ctx) error {
	result := u.Facade.TestValidate(&entity.User{})
	return c.Status(fiber.StatusInternalServerError).JSON(result)
}

func (u *UserController) MapRequestContext(fun func(ctx *fiber.Ctx) error) {

}
