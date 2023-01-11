package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ShowCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	code := ctx.Params("code")

	categoryData, err := category.Repository().Get(code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter categoria (%v).", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS055"))
	}

	if categoryData == nil {
		log.Error().Msgf("Catgeoria não encontrada (%v).", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS198", "Essa categoria não foi encontrada na base de dados."))
	}

	return response.Ctx(ctx).Result(response.Success(200, categoryData))

}
