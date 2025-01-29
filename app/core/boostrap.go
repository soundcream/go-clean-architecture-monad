package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"log"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/validators"
	"os"
)

type Application struct {
	app    *fiber.App
	Config *global.Config
}

func (a *Application) SetupLog() {
	// Initialize default config
	a.app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	a.app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func (a *Application) SetupSwagger() {
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
	a.app.Get("/swagger/*", swagger.New(swaggerConf))
}

func (a *Application) SetupI18n() {
	a.app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:         "./i18n",
			AcceptLanguages:  []language.Tag{language.Thai, language.English},
			DefaultLanguage:  language.Thai,
			FormatBundleFile: "yaml",
		}),
	)
	a.app.Get("/hello", func(c *fiber.Ctx) error {
		localize, err := fiberi18n.Localize(c, "welcome")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(localize)
	})
	a.app.Get("/hello/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fiberi18n.MustLocalize(ctx, &i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": ctx.Params("name"),
			},
		}))
	})
}

func (a *Application) SetupAppConfig() {
	//if err != nil {
	//	panic(fmt.Errorf("fatal error config file: %w", err))
	//}
	env := os.Getenv(base.Environment)
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// # WatchConfig
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	// # WatchConfig

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	a.Config = &global.Config{}
	err = viper.Unmarshal(a.Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	appName := viper.GetString("app.appName")
	domain := viper.GetString("app.domain")
	fmt.Printf("Initialized (%s)Config... App:%s; Domain:%s", env, appName, domain)
}

func (a *Application) SetupValidator() {
	myValidator := &global.XValidator{
		Validator: validator.New(),
	}
	err := validators.RegisterIsTeenValidator(myValidator.Validator)
	if err != nil {
		return
	}
	a.app.Get("/validate", func(c *fiber.Ctx) error {
		user := &global.User{
			Name: "",
			Age:  0,
		}
		errs := myValidator.Validate(user)
		result := base.NewErrContextFromInvalidateField(errs)
		return c.JSON(result)
	})
}

func (a *Application) Bootstrapper() {
	a.SetupAppConfig()
	a.SetupLog()
	a.SetupI18n()
	a.SetupValidator()
	a.SetupSwagger()
	a.MapRoute()
}

func (a *Application) Start() {
	a.app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Service-Header", fmt.Sprintf("%s:%s", a.Config.App.AppName, a.Config.App.Domain))
		return c.Next()
	})

	// Start the server
	port := a.Config.App.HttpPort
	err := a.app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewApp() Application {
	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("%+v", err)
			return c.Status(fiber.StatusBadRequest).JSON(global.ErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	a := Application{
		app: app,
	}
	return a
}
