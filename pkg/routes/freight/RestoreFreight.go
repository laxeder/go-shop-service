package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func RestoreFreight(ctx *fiber.Ctx) error {

	var log = logger.New()

	uid := ctx.Params("uid")

	freightDatabase, err := freight.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar frete. (%v)", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS070"))
	}

	// verifica o status do frete
	if freightDatabase.Status == freight.Enabled {
		log.Error().Msgf("Este frete já está ativada no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS071", "Este frete já está ativo no sistema."))
	}

	freightDatabase.Uid = uid
	freightDatabase.Status = freight.Enabled
	freightDatabase.UpdatedAt = date.NowUTC()

	err = freight.Repository().Restore(freightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS072", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
