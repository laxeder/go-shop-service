package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
)

func RestoreShopCart(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	shopcartDatabase, err := shopcart.Repository().GetUuid(uuid)
	fmt.Print(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar shopcart. (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS104"))
	}

	if shopcartDatabase.Status != shopcart.Disabled {
		log.Error().Msgf("Este shopcart já está ativo no sistema. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS105", "Este shopcart já está ativo no sistema."))
	}

	shopcartDatabase.Uuid = uuid
	shopcartDatabase.Status = shopcart.Enabled
	shopcartDatabase.LastAcesses = date.NowUTC()

	// salva as alterações na base de dados
	err = shopcart.Repository().Restore(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS106", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
