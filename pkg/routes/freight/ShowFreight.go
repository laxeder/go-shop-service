package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ShowFreight(ctx *fiber.Ctx) error {

	var log = logger.New()

	uid := ctx.Params("uid")

	freightData, err := freight.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do frete %v", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS130"))
	}

	return response.Ctx(ctx).Result(response.Success(200, freightData))

}
