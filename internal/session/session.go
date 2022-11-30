package session

import (
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis"
	"github.com/spf13/viper"
)

var sessionStore session.Store

// Init 初始化 session 组件
func Init() {
	sessionConfig := session.Config{
		Expiration:   24 * time.Hour,
		KeyLookup:    "cookie:session_id",
		KeyGenerator: utils.UUIDv4,
	}

	// 如果配置中有 redis 的配置，则使用 redis 存储 session 信息
	if viper.InConfig("db.redis") {
		// Initialize custom config
		store := redis.New(redis.Config{
			Host:      viper.GetString("db.redis.host"),
			Port:      viper.GetInt("db.redis.port"),
			Username:  viper.GetString("db.redis.username"),
			Password:  viper.GetString("db.redis.password"),
			Database:  viper.GetInt("db.redis.database"),
			Reset:     false,
			TLSConfig: nil,
			PoolSize:  10 * runtime.GOMAXPROCS(0),
		})
		sessionConfig.Storage = store
	}

	sessionStore = *session.New(sessionConfig)
}

// Get 获取 session
func Get(c *fiber.Ctx) (*session.Session, error) {
	return sessionStore.Get(c)
}
