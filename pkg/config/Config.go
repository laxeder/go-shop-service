package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

func Server() *fiber.Config {

	var log = logger.New()

	readTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Error().Err(err).Msg("Não foi possível definir uma timeout para a API")
	}

	appName := os.Getenv("APPNAME")

	return &fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		AppName:     appName,
	}
}
