package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func CreateShopCart(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	shopcartBody, err := shopcart.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos ou json está mal formatado. %s", shopcartBody)
		return response.Ctx(ctx).Result(response.Error(400, "GSS095", "Os campos enviados estão incorretos ou json está mal formatado."))
	}

	shopcartDatabase, err := shopcart.Repository().GetUuid(shopcartBody.Uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos (%v). %v", shopcartBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS096"))
	}

	if shopcartDatabase.Status == shopcart.Disabled {
		log.Error().Msgf("Este shopcart (%v) está desabilitada por tempo indeterminado.", shopcartBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS097", "Este shopcart está desabilitada por tempo indeterminado."))
	}

	if len(shopcartDatabase.Uuid) > 0 {
		log.Error().Msgf("Este shopcart (%v) já existe na nossa base de dados.", shopcartBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS098", "Este shopcart já existe na nossa base de dados."))
	}

	shopcartBody.Status = shopcart.Enabled
	shopcartBody.LastAcesses = date.NowUTC()

	err = shopcart.Repository().Save(shopcartBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do carrinho de compras %v", shopcartBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS099"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
