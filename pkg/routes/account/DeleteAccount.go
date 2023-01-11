package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func DeleteAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	accountData, err := account.Repository().GetDataInfo(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter conta (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS011"))
	}

	if accountData == nil {
		log.Error().Msgf("Conta não encontrada. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS012", "Essa conta não foi encontrada na base de dados."))
	}

	if accountData.Status != status.Enabled {
		log.Error().Msgf("Conta já está desativada (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS194", "Esta conta já está desativada no sistema."))
	}

	err = account.Repository().Delete(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar deletar conta (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS013"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
