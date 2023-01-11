package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func DeleteShopCart(ctx *fiber.Ctx) error {
	var log = logger.New()

	uuid := ctx.Params("uuid")

	shopcartDatabase, err := shopcart.Repository().GetUuid(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS100"))
	}

	if shopcartDatabase.Status != shopcart.Enabled {
		log.Error().Msgf("Este shopcart já está desativado no sistema. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS101", "Este shopcart já está desativado no sistema."))
	}

	shopcartDatabase.Uuid = uuid
	shopcartDatabase.Status = shopcart.Disabled
	shopcartDatabase.LastAcesses = date.NowUTC()

	err = shopcart.Repository().Delete(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS102"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
