package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/utils"
	"log"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/validators"
	_ "n4a3/clean-architecture/docs"
)

func main() {
	app := core.NewApp()

	core.SetupLog(app)
	core.SetupSwagger(app)

	myValidator := &global.XValidator{
		Validator: validator.New(),
	}

	//app.Use(
	//	fiberi18n.New(&fiberi18n.Config{
	//		RootPath:        "./example/localize",
	//		AcceptLanguages: []language.Tag{language.Chinese, language.English},
	//		DefaultLanguage: language.Chinese,
	//	}),
	//)

	// Validator
	err := validators.RegisterIsTeenValidator(myValidator.Validator)
	if err != nil {
		return
	}
	app.Get("/validate", func(c *fiber.Ctx) error {
		user := &global.User{
			Name: "",
			Age:  0,
		}
		errs := myValidator.Validate(user)
		result := base.NewErrContextFromInvalidateField(errs)
		return c.JSON(result)
	})

	core.MapRoute(app)

	// set header
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "EX_HEADER")
		return c.Next()
	})

	// Start the server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
