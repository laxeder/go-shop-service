package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func UpdateFreight(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	uid := ctx.Params("uid")

	freightBody, err := freight.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS074", "O formado dos dados envidados está incorreto."))
	}

	freightDatabase, err := freight.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar frete %v.", freightBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS075", "Erro ao tentar validar frete."))
	}

	freightDatabase.Inject(freightBody)
	freightDatabase.UpdatedAt = date.NowUTC()

	err = freight.Repository().Update(freightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar o frete %v no repositório", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS076"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
