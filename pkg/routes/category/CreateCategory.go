package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func CreateCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	categoryBody := &category.Category{}

	err := utils.InjectBytes(body, categoryBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body na categoria. %s", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS039", "O formado dos dados envidados está incorreto."))
	}

	categoryData, err := category.Repository().Get(categoryBody.Code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter categoria pelo code (%v)", categoryBody.Code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS040"))
	}

	if categoryData != nil {
		log.Error().Msgf("Categoria já registrada.", categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS041", "Essa categoria já está registrada na base de dados."))
	}

	err = category.Repository().Save(categoryBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar salvar categoria (%v)", categoryBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS043"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
