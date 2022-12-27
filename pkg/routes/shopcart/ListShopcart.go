package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/shopcart"
)

func ListShopCarts(ctx *fiber.Ctx) error {

	var log = logger.New()

	shopcarts, err := shopcart.Repository().GetList()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar carrinhos de compras. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS103"))
	}

	return response.Ctx(ctx).Result(response.Success(200, shopcarts))

}
