package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func DeleteCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	code := ctx.Params("code")

	categoryData, err := category.Repository().GetDataInfo(code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter categoria (%v).", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS044"))
	}

	if categoryData == nil {
		log.Error().Msgf("Categoria não encontrada (%v).", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS185", "Essa categoria não foi encontrada na base de dados."))
	}

	if categoryData.Status != status.Enabled {
		log.Error().Msgf("Categoria já está desativada. (%v)", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS045", "Está categoria já está desativada no sistema."))
	}

	err = category.Repository().Delete(code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar deletar categoria (%v).", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS046"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
