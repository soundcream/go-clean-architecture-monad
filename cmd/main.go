package main

import (
	_ "github.com/gofiber/fiber/v2/utils"
	"n4a3/clean-architecture/app/http-api"
	_ "n4a3/clean-architecture/docs"
)

// @title Swagger API
// @version 1.0
// @description Doc
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app := http_api.NewApp()
	app.Bootstrapper()
	app.Start()
}
