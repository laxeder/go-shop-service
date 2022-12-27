package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// restaura um categoria com status deletado
func RestoreCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	code := ctx.Params("code")

	// carrega a categoria com base no code
	categoryDatabase, err := category.Repository().GetByCode(code)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar categoria. (%v)", code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS048"))
	}

	// verifica o status da categoria
	if categoryDatabase.Status != category.Disabled {
		log.Error().Msgf("Este categoria já está ativo no sistema. (%v)", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS049", "Este categoria já está ativo no sistema."))
	}

	// muda o status da categoria para ativo
	categoryDatabase.Status = category.Enabled
	categoryDatabase.UpdatedAt = date.NowUTC()
	categoryDatabase.Code = code

	// salva as alterações na base de dados
	err = category.Repository().Restore(categoryDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS050", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
