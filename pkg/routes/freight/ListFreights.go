package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListFreights(ctx *fiber.Ctx) error {

	var log = logger.New()

	freights, err := freight.Repository().GetList()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar fretes %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS069"))
	}

	return response.Ctx(ctx).Result(response.Success(200, freights))

}
