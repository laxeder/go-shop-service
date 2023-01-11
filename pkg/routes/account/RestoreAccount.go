package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func RestoreAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	accountData, err := account.Repository().GetDataInfo(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter conta. (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS015"))
	}

	if accountData == nil {
		log.Error().Msgf("Conta já está ativada (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS195", "Essa conta não foi encontrada na base de dados."))
	}

	if accountData.Status != status.Disabled {
		log.Error().Msgf("Conta já está ativada. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS016", "Este conta já está ativo no sistema."))
	}

	err = account.Repository().Restore(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar restaurar conta. (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS126"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
