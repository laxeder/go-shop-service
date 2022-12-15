package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// restaura um produto com status deletado
func RestoreProduct(ctx *fiber.Ctx) error {

	var log = logger.New()

	uid := ctx.Params("uid")

	// carrega o produto com base no uid
	productDatabase, err := product.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar produto. (%v)", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC097"))
	}

	// verifica o status do produto
	if productDatabase.Status != product.Disabled {
		log.Error().Msgf("Este produto já está ativo no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este produto já está ativo no sistema."))
	}

	// muda o status do produto para ativo
	productDatabase.Status = product.Enabled
	productDatabase.UpdatedAt = date.NowUTC()
	productDatabase.Uid = uid

	// salva as alterações na base de dados
	err = product.Repository().Restore(productDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC100", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
