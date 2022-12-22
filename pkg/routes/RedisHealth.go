package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// retorna a hora do banco de dados
func RedisHealth(ctx *fiber.Ctx) error {
	var log = logger.New()

	result, err := redisdb.Health()
	if err != nil {
		log.Error().Err(err).Msg("O banco de dados está offline")
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS002"))
	}

	return response.Ctx(ctx).Result(response.Success(200, result))
}
