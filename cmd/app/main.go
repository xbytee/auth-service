package main

import (
	"auth-service/config"
	"auth-service/pkg/logger"
	"fmt"
)

func main() {
	cfg := config.ReadConfig("./config/config.yaml")

	l := logger.New(cfg.App.Logger.Mode)

	fmt.Println(cfg, l)
}
