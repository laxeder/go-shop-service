package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
)

// mostra os dados de uma usu'ario
func ShowShopCart(ctx *fiber.Ctx) error {
	var log = logger.New()

	uuid := ctx.Params("uuid")

	shopcartDatabase, err := shopcart.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS107"))
	}

	return response.Ctx(ctx).Result(response.Success(200, shopcartDatabase))

}
