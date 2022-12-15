package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// muda o status do produto na base de dados
func DeleteProduct(ctx *fiber.Ctx) error {
	var log = logger.New()

	uid := ctx.Params("uid")

	productDatabase, err := product.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC087"))
	}

	// verifica o status do produto
	if productDatabase.Status != product.Enabled {
		log.Error().Msgf("Este produto já está desativado no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este produto já está desativado no sistema."))
	}

	productDatabase.Status = product.Disabled
	productDatabase.UpdatedAt = date.NowUTC()
	productDatabase.Uid = uid

	err = product.Repository().Delete(productDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
