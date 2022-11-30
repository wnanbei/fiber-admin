package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	localLogger "github.com/wnanbei/fiber-admin/internal/logger"
)

// setGlobalMiddlewares 设置全局中间件
func setGlobalMiddlewares(app *fiber.App) {
	// logger config
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}\n",
		TimeFormat:   "2006-01-02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       localLogger.Writer,
	}

	// limiter
	limiterConfig := limiter.Config{
		Max:        viper.GetInt("server.globalLimiterMax"),
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
