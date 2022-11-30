package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	"github.com/wnanbei/fiber-admin/controller/ping"
	localLogger "github.com/wnanbei/fiber-admin/internal/logger"
	"github.com/wnanbei/fiber-admin/internal/session"
	swag "github.com/wnanbei/fiber-admin/internal/swagger"
)

// @title Fiber Example API
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func New() {
	app := fiber.New()
	setGlobalMiddlewares(app)

	if viper.GetBool("server.enableSwagger") {
		setSwagger(app)
	}

	setRouter(app)

	session.Init()

	app.Listen(":" + viper.GetString("server.port"))
}

func setRouter(app *fiber.App) {
	app.Get("/ping", ping.Ping)
}

// SetSwagger 设置 swagger 文档
func setSwagger(app *fiber.App) {
	swag.SwaggerInfo.Title = viper.GetString("server.title")
	swag.SwaggerInfo.Host = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	swag.SwaggerInfo.BasePath = viper.GetString("server.basePath")
	swag.SwaggerInfo.Version = viper.GetString("server.version")
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}

// setGlobalMiddlewares 设置全局中间件
func setGlobalMiddlewares(app *fiber.App) {
	// logger config
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       localLogger.Writer,
	}

	// limiter
	limiterConfig := limiter.Config{
		Max:        viper.GetInt("server.limiterMax"),
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
	}

	app.Use(
		logger.New(loggerConfig),
		limiter.New(limiterConfig),
		requestid.New(),
	)
}
