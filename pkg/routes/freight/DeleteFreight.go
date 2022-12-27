package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func DeleteFreight(ctx *fiber.Ctx) error {
	var log = logger.New()

	uid := ctx.Params("uid")

	freightDatabase, err := freight.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS087"))
	}

	if freightDatabase.Status == freight.Disabled {
		log.Error().Msgf("Este frete já está desativado no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS060", "Este frete já está desativado no sistema."))
	}

	freightDatabase.Uid = uid
	freightDatabase.Status = freight.Disabled
	freightDatabase.UpdatedAt = date.NowUTC()

	err = freight.Repository().Delete(freightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
