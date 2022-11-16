package main

import (
	"github.com/wnanbei/fiber-admin/internal/config"
	"github.com/wnanbei/fiber-admin/internal/router"
)

func init() {}

func main() {
	config.Init()

	router.New()
}
