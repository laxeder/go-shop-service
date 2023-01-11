package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func RestoreCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	code := ctx.Params("code")

	categoryData, err := category.Repository().GetDataInfo(code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter categoria. (%v)", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS048"))
	}

	if categoryData == nil {
		log.Error().Msgf("Categoria não encontrada (%v).", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS197", "Essa categoria não foi encontrada na base de dados."))
	}

	if categoryData.Status != status.Disabled {
		log.Error().Msgf("Usuário já está ativado. (%v)", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS049", "Essa categoria já está ativada na base de dados."))
	}

	err = category.Repository().Restore(code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar restaurar categoria. (%v)", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS050"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
