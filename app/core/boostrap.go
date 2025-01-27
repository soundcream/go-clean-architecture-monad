package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterLog(app *fiber.App) {
	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}
