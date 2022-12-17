package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// mostra os dados de uma categoria
func ShowCategory(ctx *fiber.Ctx) error {
	var log = logger.New()

	code := ctx.Params("code")

	// carrega uma categoria da base de dados
	categoryData, err := category.Repository().GetByCode(code)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS035"))
	}

	return response.Ctx(ctx).Result(response.Success(200, categoryData))

}
