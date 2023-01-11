package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	categoryBody := &category.Category{}

	err := utils.InjectBytes(body, categoryBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body na categoria. (%s)", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS056", "O formado dos dados envidados está incorreto."))
	}

	categoryData, err := category.Repository().Get(categoryBody.Code)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter a categoria (%v).", categoryBody.Code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS058"))
	}

	if categoryData == nil {
		log.Error().Msgf("Categoria não encontrada (%v).", categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS057", "Essa categoria não foi encontrada na base de dados."))
	}

	utils.Inject(categoryBody, categoryData)

	err = category.Repository().Update(categoryData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar categoria (%v).", categoryBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS059"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
