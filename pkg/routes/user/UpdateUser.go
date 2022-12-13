package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func UpdateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// converte json para struct
	userBody, err := user.New(body)
	if err != nil {
		log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
		return response.Ctx(ctx).Result(response.Error(400, "BLC085", "O formado dos dados envidados está incorreto."))
	}

	// verifica e compara o documento recebido
	if userBody.Document != "" && str.DocumentClean(document) != str.DocumentClean(userBody.Document) {
		log.Error().Msgf("Não é possível atualiztar o documento %v para o %v", document, userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC086", "Não é possível atualiztar do documento "+document+" para o "+userBody.Document))
	}

	// carrega o usuário da base de dados
	userDatabase, err := user.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário %v.", userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC081", "Erro ao tentar validar usuário."))
	}

	// injecta dos dados novos o lugar dos dsdos trazidos d abase de dados
	userDatabase.Inject(userBody)
	userDatabase.UpdatedAt = date.NowUTC()
	userDatabase.Document = document
	userDatabase.SetFullname()

	// guarda as alterações do usuário na base de dados
	err = user.Repository().Update(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do usuário (%v)", userBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC084"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
