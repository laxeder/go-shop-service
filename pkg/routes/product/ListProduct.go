package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListProducts(ctx *fiber.Ctx) error {

	var log = logger.New()

	products, err := product.Repository().GetList()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar produtos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS085"))
	}

	return response.Ctx(ctx).Result(response.Success(200, products))

}
