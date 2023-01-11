package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func UpdateShopCart(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	uuid := ctx.Params("uuid")

	shopcartBody, err := shopcart.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS108", "O formado dos dados envidados está incorreto."))
	}

	shopcartDatabase, err := shopcart.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS109"))
	}

	shopcartDatabase.Inject(shopcartBody)
	shopcartDatabase.LastAcesses = date.NowUTC()

	// guarda as alterações do shopcart na base de dados
	err = shopcart.Repository().Update(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do shopcart (%v)", shopcartBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS110"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
