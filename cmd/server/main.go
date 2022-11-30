package main

import (
	"github.com/wnanbei/fiber-admin/internal/config"
	"github.com/wnanbei/fiber-admin/internal/logger"
	"github.com/wnanbei/fiber-admin/router"
)

func init() {}

func main() {
	config.Init()
	logger.Init()
	// mysql.Init()
	router.New()
}
