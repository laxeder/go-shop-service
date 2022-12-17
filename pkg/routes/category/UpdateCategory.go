package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateCategory(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	code := ctx.Params("code")

	// converte json para struct
	categoryBody, err := category.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS085", "O formado dos dados envidados está incorreto."))
	}

	// verifica e compara a categoria recebida
	if categoryBody.Code != "" && code != categoryBody.Code {
		log.Error().Msgf("Não é possível atualizar a categoria %v para o %v", code, categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS086", "Não é possível atualizar a categoria "+code+" para o "+categoryBody.Code))
	}

	// carrega a categoria da base de dados
	categoryDatabase, err := category.Repository().GetByCode(code)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar a categoria %v.", categoryBody.Code)
		return response.Ctx(ctx).Result(response.Error(400, "GSS081", "Erro ao tentar validar a categoria."))
	}

	// injeta os dados novos no lugar dos dados trazidos da base de dados
	categoryDatabase.Inject(categoryBody)
	categoryDatabase.UpdatedAt = date.NowUTC()
	categoryDatabase.Code = code

	// guarda as alterações da categoria na base de dados
	err = category.Repository().Update(categoryDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório da categoria (%v)", categoryBody.Code)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS084"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
