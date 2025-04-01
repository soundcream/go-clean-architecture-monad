package controllers

import (
	"github.com/gofiber/fiber/v2"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
	"n4a3/clean-architecture/app/core/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/integrates/cache"
	"n4a3/clean-architecture/app/integrates/db"
	"n4a3/clean-architecture/app/integrates/dto"
	"n4a3/clean-architecture/app/integrates/repository"
	"n4a3/clean-architecture/app/integrates/websockets"
)

type DemoController struct {
	Config        *global.Config
	WS            *websockets.WebSocketServer
	RedisDb       cache.RedisContext
	Facade        facades.DemoFacade
	QueryFacade   facades.QueryFacade
	CommandFacade facades.CommandFacade
}

func NewDemoController(config *global.Config, ws *websockets.WebSocketServer) *DemoController {
	cmdUoW := db.NewUnitOfWork(config)
	queryUoW := db.NewQueryUnitOfWork(config)
	repo := repository.NewUserRepository(queryUoW.Right, cmdUoW.Right)
	redisDb := cache.NewRedisContext(config.RedisConfig)
	return &DemoController{
		Config:        config,
		WS:            ws,
		RedisDb:       redisDb,
		Facade:        facades.NewDemoFacade(*config),
		QueryFacade:   facades.NewQueryFacade(repo),
		CommandFacade: facades.NewCommandFacade(repo),
	}
}

func (con *DemoController) MapRoute(route fiber.Router) {
	route.Post("/ex", func(c *fiber.Ctx) error {
		return con.TestValidate(c)
	})
	route.Get("/user", func(c *fiber.Ctx) error {
		return con.GetUserById(c)
	})
	route.Get("/users", func(c *fiber.Ctx) error {
		return con.SearchUsers(c)
	})
	route.Post("/insert", func(c *fiber.Ctx) error {
		return con.Insert(c)
	})
	route.Put("/:id/update", func(c *fiber.Ctx) error {
		return con.Update(c)
	})
	route.Delete("/:id/delete", func(c *fiber.Ctx) error {
		return con.Delete(c)
	})
	route.Post("/:id/updatew", func(c *fiber.Ctx) error {
		return con.UpdateWhere(c)
	})
	route.Post("/:id/updatefw", func(c *fiber.Ctx) error {
		return con.UpdateFieldWhere(c)
	})
	route.Get("/test-http", func(c *fiber.Ctx) error {
		return con.RequestHttp(c)
	})
	route.Post("/mapper", func(c *fiber.Ctx) error {
		return con.TestMap(c)
	})
	route.Post("/ws/cmd", func(c *fiber.Ctx) error {
		return con.WsCmd(c)
	})
	route.Get("/cache/get", func(c *fiber.Ctx) error {
		return con.CacheGet(c)
	})
	route.Get("/cache/set", func(c *fiber.Ctx) error {
		return con.CacheSet(c)
	})
	route.Post("/cache/pub", func(c *fiber.Ctx) error {
		return con.CachePub(c)
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

func (con *DemoController) TestMap(c *fiber.Ctx) error {
	result := con.QueryFacade.GetUser()
	userCreate := "Z"
	var u1 = entity.User{
		BaseEntity: &entity.BaseEntity{
			Id:        10,
			CreatedBy: userCreate,
		},
		Name: "A",
	}
	var u2 = entity.User{}
	util.MapValue(&u1, &u2)
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}

func (con *DemoController) CacheGet(c *fiber.Ctx) error {
	result := con.RedisDb.GetWithKey("name_key")
	return Response(c, result)
}

func (con *DemoController) CacheSet(c *fiber.Ctx) error {
	v := c.Query("v", "")
	result := con.RedisDb.SetWithKey("name_key", v)
	return Response(c, result)
}

func (con *DemoController) CachePub(c *fiber.Ctx) error {
	v := c.Query("m", "")
	result := con.RedisDb.Publish("ws-msg", v)
	return Response(c, result)
}

func (con *DemoController) WsCmd(c *fiber.Ctx) error {
	if con.WS != nil {
		(*con.WS).BroadcastCmd(websockets.WsCommand{
			Msg:     "FromDemoController",
			Command: "Msg",
		})
	}
	return Response(c, core.RightEither[core.Unit, core.ErrContext](core.Unit{}))
}

// GetUserById @Summary Example GetUser
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/user [get]
func (con *DemoController) GetUserById(c *fiber.Ctx) error {
	result := con.QueryFacade.GetUserById(c.QueryInt("id"))
	return Response(c, result)
}

// SearchUsers @Summary Example SearchUsers
func (con *DemoController) SearchUsers(c *fiber.Ctx) error {
	result := con.QueryFacade.SearchUsers(
		c.Query("keyword", ""),
		c.QueryInt("limit", 10),
		c.QueryInt("offset", 0))
	return Response(c, result)
}

// Insert @Summary Example Insert
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/insert [post]
func (con *DemoController) Insert(c *fiber.Ctx) error {
	return Response(c, MapBody[dto.UserDto](c).
		Then(con.CommandFacade.Insert))
}

// Update @Summary Example Update
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/:id/update [post]
func (con *DemoController) Update(c *fiber.Ctx) error {
	return Response(c, MapCommandByRouteParamsId[dto.UserDto](c).
		Then(con.CommandFacade.Update))
}

// Delete @Summary Example Delete
// @Description
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /api/demo/insert [post]
func (con *DemoController) Delete(c *fiber.Ctx) error {
	return Response(c, GetRouteParamsById(c).
		Then(con.CommandFacade.Delete))
}

func (con *DemoController) UpdateWhere(c *fiber.Ctx) error {
	return Response(c, GetRouteParamsById(c).
		Then(con.CommandFacade.UpdateWhere))
}

func (con *DemoController) UpdateFieldWhere(c *fiber.Ctx) error {
	return Response(c, GetRouteParamsById(c).
		Then(con.CommandFacade.UpdateFieldWhere))
}

func (con *DemoController) RequestHttp(c *fiber.Ctx) error {
	return Response(c, con.Facade.RequestHttp())
}

func (con *DemoController) Insert1(c *fiber.Ctx) error {
	u := new(dto.UserDto)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	input := util.MapFrom[dto.UserDto](u)
	result := con.CommandFacade.Insert(*input)
	if result.IsRight() {
		return OkResult(c, result.Right)
	}
	return ErrorResult(c, result.Left)
}
