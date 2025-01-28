package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"n4a3/clean-architecture/app/base/global"
)

func SetupLog(app *fiber.App) {
	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func SetupSwagger(app *fiber.App) {

	swaggerConf := swagger.Config{
		Title: "",
		//URL:          "/docs/swagger.json",
		DeepLinking:              true,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DocExpansion:             "list",
		OAuth: &swagger.OAuthConfig{
			AppName:                           "OAuth Provider",
			ClientId:                          "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
			UsePkceWithAuthorizationCodeGrant: true,
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}

	//app.Get("/swagger/*", swagger.HandlerDefault) // Serve Swagger UI
	app.Get("/swagger/*", swagger.New(swaggerConf))
}

func NewApp() *fiber.App {
	//app := fiber.New()
	return fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(global.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
}
