package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	addresss, err := address.Repository().GetList()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar endereços, %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS031"))
	}

	return response.Ctx(ctx).Result(response.Success(200, addresss))

}
