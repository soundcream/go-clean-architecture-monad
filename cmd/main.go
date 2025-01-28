package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/utils"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
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

	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:         "./i18n",
			AcceptLanguages:  []language.Tag{language.Thai, language.English},
			DefaultLanguage:  language.Thai,
			FormatBundleFile: "yaml",
		}),
	)

	app.Get("/hello", func(c *fiber.Ctx) error {
		localize, err := fiberi18n.Localize(c, "welcome")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(localize)
	})
	app.Get("/hello/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fiberi18n.MustLocalize(ctx, &i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": ctx.Params("name"),
			},
		}))
	})

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
