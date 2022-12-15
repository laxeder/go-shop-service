package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// mostra os dados de um endereço
func ShowAddress(ctx *fiber.Ctx) error {
	var log = logger.New()

	uid := ctx.Params("uid")

	addressData, err := address.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos., %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC035"))
	}

	return response.Ctx(ctx).Result(response.Success(200, addressData))

}
