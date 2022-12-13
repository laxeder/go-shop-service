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

	document := ctx.Params("document")

	// carega um endereço da base de dados
	addressData, err := address.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msg("Os campos enviados estão incorretos.")
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC035"))
	}

	return response.Ctx(ctx).Result(response.Success(200, addressData))

}
