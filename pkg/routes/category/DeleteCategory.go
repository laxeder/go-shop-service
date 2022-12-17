package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// muda o status da categoria na base de dados
func DeleteCategory(ctx *fiber.Ctx) error {
	var log = logger.New()

	code := ctx.Params("code")

	categoryDatabase, err := category.Repository().GetByCode(code)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS087"))
	}

	// verifica o status da categoria
	if categoryDatabase.Status != category.Enabled {
		log.Error().Msgf("Está categoria já está desativado no sistema. (%v)", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS060", "Está categoria já está desativado no sistema."))
	}

	categoryDatabase.Status = category.Disabled
	categoryDatabase.UpdatedAt = date.NowUTC()
	categoryDatabase.Code = code

	err = category.Repository().Delete(categoryDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
