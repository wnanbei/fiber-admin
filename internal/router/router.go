package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	localLogger "github.com/wnanbei/fiber-admin/internal/logger"
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
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       localLogger.Writer,
	}

	app := fiber.New()
	app.Use(
		logger.New(loggerConfig),
		requestid.New(),
	)

	if viper.GetBool("server.enableSwagger") {
		SetSwagger(app)
	}

	SetRouter(app)

	app.Listen(":" + viper.GetString("server.port"))
}

func SetRouter(app *fiber.App) {}

// SetSwagger 设置 swagger 文档
func SetSwagger(app *fiber.App) {
	swag.SwaggerInfo.Title = viper.GetString("server.title")
	swag.SwaggerInfo.Host = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	swag.SwaggerInfo.BasePath = viper.GetString("server.basePath")
	swag.SwaggerInfo.Version = viper.GetString("server.version")
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}
