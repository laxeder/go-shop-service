package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ShowAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	accountData, err := account.Repository().Get(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter conta (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS018"))
	}

	if accountData == nil {
		log.Error().Msgf("Conta não encontrada (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS196", "Essa conta não foi encontrada na base de dados."))
	}

	return response.Ctx(ctx).Result(response.Success(200, accountData))

}
