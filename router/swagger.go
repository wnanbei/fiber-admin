package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	swag "github.com/wnanbei/fiber-admin/internal/swagger"
)

// SetSwagger 设置 swagger 文档
func setSwagger(app *fiber.App) {
	swag.SwaggerInfo.Title = viper.GetString("server.title")
	swag.SwaggerInfo.Host = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	swag.SwaggerInfo.BasePath = viper.GetString("server.basePath")
	swag.SwaggerInfo.Version = viper.GetString("server.version")
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}
