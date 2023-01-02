package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
)

func RemoveShopCartProduct(ctx *fiber.Ctx) error {

	var log = logger.New()

	productsBody := ctx.Body()
	uuid := ctx.Params("uuid")

	shopcartDatabase, err := shopcart.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS109"))
	}

	var products []product.Product

	products = shopcart.UnmarshalBinary(productsBody).([]product.Product)

	shopcartDatabase.RemoveLote(products)
	shopcartDatabase.LastAcesses = date.NowUTC()

	// guarda as alterações do shopcart na base de dados
	err = shopcart.Repository().Update(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do shopcart (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS110"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
