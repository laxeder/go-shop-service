package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func AddShopCartProduct(ctx *fiber.Ctx) error {

	var log = logger.New()

	productsBody := ctx.Body()
	uuid := ctx.Params("uuid")

	shopcartDatabase, err := shopcart.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS109"))
	}

	var products []product.Product

	err = json.Unmarshal(productsBody, &products)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar bytes em array %s", productsBody)
	}

	shopcartDatabase.Lote(products)
	shopcartDatabase.LastAcesses = date.NowUTC()

	// guarda as alterações do shopcart na base de dados
	err = shopcart.Repository().Update(shopcartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do shopcart (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS110"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
