package router

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"github.com/wnanbei/fiber-admin/controller/ping"
	"github.com/wnanbei/fiber-admin/internal/session"
)

//	@title			Fiber Example API
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func New() {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: jsoniter.ConfigCompatibleWithStandardLibrary.Marshal,
			JSONDecoder: jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
		},
	)
	setGlobalMiddlewares(app)
	setRouter(app)
	session.Init()
	app.Listen(":" + viper.GetString("server.port"))
}

func setRouter(app *fiber.App) {
	if viper.GetBool("server.enableSwagger") {
		setSwagger(app)
	}

	app.Get("/ping", ping.Ping)
}
