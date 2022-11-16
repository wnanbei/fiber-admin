package router

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func New() {
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}

	app := fiber.New()
	app.Use(
		logger.New(loggerConfig),
	)

	Router(app)

	app.Listen(":" + viper.GetString("server.port"))
}

func Router(app *fiber.App) {}
