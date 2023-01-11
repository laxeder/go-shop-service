package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ShowAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")
	uid := ctx.Params("uid")

	addressData, err := address.Repository().Get(uuid, uid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter endereço (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS035"))
	}

	if addressData == nil {
		log.Error().Msgf("Endereço não encontrado (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS201", "Esse endereço não foi encontrado na base de dados."))
	}

	return response.Ctx(ctx).Result(response.Success(200, addressData))

}
