package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListCategories(ctx *fiber.Ctx) error {

	var log = logger.New()

	categories, err := category.Repository().GetList()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar categorias. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS047"))
	}

	return response.Ctx(ctx).Result(response.Success(200, categories))

}
