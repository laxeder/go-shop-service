package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

func UpdateAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// converte json para struct
	addressBody, err := address.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "BLC085", "O formado dos dados envidados está incorreto."))
	}

	// verifica e compara o documento recebido
	if addressBody.Document != "" && str.DocumentClean(document) != str.DocumentClean(addressBody.Document) {
		log.Error().Msgf("Não é possível atualiztar o documento %v para o %v", document, addressBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC086", "Não é possível atualiztar do documento "+document+" para o "+addressBody.Document))
	}

	// carrega o endereço da base de dados
	addressDatabase, err := address.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar endereço %v.", addressBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC081", "Erro ao tentar validar endereço."))
	}

	// injecta dos dados novos o lugar dos dsdos trazidos d abase de dados
	addressDatabase.Inject(addressBody)
	addressDatabase.UpdatedAt = date.NowUTC()
	addressDatabase.Document = document

	// guarda as alterações do endereço na base de dados
	err = address.Repository().Update(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do endereço (%v)", addressBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC084"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
