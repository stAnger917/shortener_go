package main

import (
	"shortener/config"
	"shortener/internal/app"
)

func main() {
	cfg := config.InitConfig("../../.env")
	app.RunApp(cfg)
}
