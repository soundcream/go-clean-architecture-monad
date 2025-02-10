package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/integrates/db"
	"n4a3/clean-architecture/app/integrates/repository"
)

type DemoController struct {
	Config        *global.Config
	Facade        facades.DemoFacade
	QueryFacade   facades.QueryFacade
	CommandFacade facades.CommandFacade
}

func NewDemoController(config *global.Config) *DemoController {
	repo := repository.NewUserRepository(db.NewQueryUnitOfWork(config).Right, db.NewUnitOfWork(config).Right)
	return &DemoController{
		Config:        config,
		Facade:        facades.NewDemoFacade(),
		QueryFacade:   facades.NewQueryFacade(repo),
		CommandFacade: facades.NewCommandFacade(repo),
	}
}

func (con *DemoController) MapRoute(route fiber.Router) {
	route.Post("/ex", func(c *fiber.Ctx) error {
		return con.TestValidate(c)
	})
	route.Get("/user", func(c *fiber.Ctx) error {
		return con.GetUser(c)
	})
	route.Post("/insert", func(c *fiber.Ctx) error {
		return con.Insert(c)
	})
}

// TestValidate @Summary Example of chain Validate
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/ex [get]
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func (con *DemoController) TestValidate(c *fiber.Ctx) error {
	result := con.Facade.Validate(new(entity.User))
	return ErrorHandleResult(c, result)
}

// GetUser @Summary Example GetUser
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/user [get]
func (con *DemoController) GetUser(c *fiber.Ctx) error {
	result := con.QueryFacade.GetUser()
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}

// Insert @Summary Example Insert
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/insert [post]
func (con *DemoController) Insert(c *fiber.Ctx) error {
	result := con.CommandFacade.Insert()
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}
