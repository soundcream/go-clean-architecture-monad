package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/core/websockets"
	"n4a3/clean-architecture/app/domain"
	"n4a3/clean-architecture/app/integrates/dto"
	"n4a3/clean-architecture/app/validators"
	"os"
	"time"
)

type AppContext struct {
	app    *fiber.App
	Config *global.Config
}

func (a *AppContext) Bootstrapper() {
	a.SetupAppConfig()
	a.SetupLog()
	a.SetupI18n()
	a.SetupValidator()
	a.SetupSwagger()
	a.SetupCustomHandler()
	a.useFavicon()
	a.SetupWebSocket()
	a.SetupAuthorization()
	a.MapRoute()
}

func (a *AppContext) SetupLog() {
	// Initialize default config
	a.app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	a.app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func (a *AppContext) SetupSwagger() {
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

func (a *AppContext) SetupI18n() {
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

func (a *AppContext) SetupAppConfig() {
	env := os.Getenv(base.Environment)
	//viper.SetConfigName(fmt.Sprintf("config.%s", env))
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./conf")
	viper.SetConfigFile(fmt.Sprintf("./conf/config.%s.yaml", env))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error read config file: %w", err))
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

func (a *AppContext) SetupValidator() {
	myValidator := &validators.XValidator{
		Validator: validator.New(),
	}
	err := validators.RegisterIsTeenValidator(myValidator.Validator)
	if err != nil {
		return
	}
	a.app.Get("/validate", func(c *fiber.Ctx) error {
		user := &domain.User{
			Name: "",
			Age:  0,
		}
		errs := myValidator.Validate(user)
		result := base.NewErrContextFromInvalidateField(errs)
		return c.JSON(result)
	})
}

// Sign-in @Summary
// @Description
// @Tags Sign-in
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/signin [get]
func (a *AppContext) SetupAuthorization() {
	// GET TOKEN
	a.app.Get("/api/signin", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"token": CreateToken()})
	})

	// Manual Authorize Handler
	//a.app.Use(jwtware.New(jwtware.Config{
	//	Filter:       SecurityFilter,
	//	ErrorHandler: UnauthorizedHandler,
	//	SigningKey:   jwtware.SigningKey{Key: []byte("secret")},
	//}))

	//a.app.Use(JwtSecurityFilter)

	// Basic Authorize Handler
	a.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Restricted Routes
	a.app.Get("/restricted", restricted)
}

func (a *AppContext) SetupCustomHandler() {
	a.app.Use(CustomHandler)
}

func (a *AppContext) SetupWebSocket() {
	server := websockets.NewWebSocket()
	a.app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))
	go server.HandleMessages()
}

func CreateToken() string {
	claims := jwt.MapClaims{
		"name":  "Arm Natthee",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Info("Error on get token %s", err)
	}
	return t
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func JwtSecurityFilter(ctx *fiber.Ctx) bool {
	name := ctx.OriginalURL()
	path := ctx.Path()
	fmt.Println(name)
	fmt.Println(path)
	return true
}

func CustomHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func UnauthorizedHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorContextResponse(base.NewErrorCode(base.Unauthorized)))
}

func (a *AppContext) useFavicon() {
	a.app.Use(favicon.New())
	// config for customization
	//a.app.Use(favicon.New(favicon.Config{
	//	File: "./favicon.ico",
	//	URL:  "/favicon.ico",
	//}))
}

func (a *AppContext) Start() {
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

func NewApp() AppContext {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Errorf("%+v", err)
			//c.OriginalURL()
			if fex, ok := err.(*fiber.Error); ok && fex.Code != 500 {
				return c.Status(fex.Code).JSON(dto.ErrorResponse(global.ErrorHandlerResp{
					Code:    fex.Code,
					Message: fex.Message,
				}))
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(global.ErrorHandlerResp{
					Code:    int(base.UnHandleError),
					Message: base.UnHandleError.GetDefaultErrorMsg(),
				}))
			}
		},
	})
	return AppContext{
		app: app,
	}
}
