package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// cria uma nova categoria na base de ddaos
func CreateCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	categoryBody, err := category.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS039", "Os campos enviados estão incorretos."))
	}

	// Cria um CODE para a categoria
	categoryBody.NewCode()

	categoryDatabase, err := category.Repository().GetByCode(categoryBody.Code)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS040"))
	}

	// verifica se a categoria está desabilitado
	if categoryDatabase.Status == category.Disabled {
		log.Error().Msgf("Está categoria (%v) está desabilitado por tempo indeterminado.", categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS041", "Está categoria está desabilitado por tempo indeterminado."))
	}

	// verifica se a categoria existe
	if len(categoryDatabase.Code) > 0 {
		log.Error().Msgf("Está categoria (%v) já existe na nossa base de dados.", categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS042", "Está categoria já existe na nossa base de dados."))
	}

	categoryBody.Status = category.Enabled
	categoryBody.CreatedAt = date.NowUTC()
	categoryBody.UpdatedAt = date.NowUTC()

	// armazena a categoria na base de dados
	err = category.Repository().Save(categoryBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório de categorias %v", categoryBody.Code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS043"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
