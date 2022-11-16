package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 初始化配置
func Init() {
	SetDefaultConfig()

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

// SetDefaultConfig 设置默认配置
func SetDefaultConfig() {
	// server
	viper.SetDefault("server.port", "5555")

	// db.mysql
	viper.SetDefault("db.mysql.username", "root")
	viper.SetDefault("db.mysql.password", "123456")
	viper.SetDefault("db.mysql.host", "127.0.0.1")
	viper.SetDefault("db.mysql.port", "3306")
	viper.SetDefault("db.mysql.dbname", "")
}
