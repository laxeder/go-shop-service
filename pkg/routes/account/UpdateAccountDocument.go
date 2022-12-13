package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateAccountDocument(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// iniicia a struct do conta com password
	accountDocument := &account.AccountDocument{}

	// converte o json para struct
	err := json.Unmarshal(body, accountDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "O formado dos dados envidados está incorreto."))
	}

	// carrega os dados do conta da base de dados
	accountDatabase, err := account.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Documento não encontrado %v.", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC093"))
	}

	if accountDocument.OldDocument != document {
		log.Error().Err(err).Msgf("Documento antigo errado %v.", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "Documento antigo está incorreto."))
	}

	// define o novo documento
	accountDatabase.SetDocument(accountDocument.Document)
	accountDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do conta
	err = account.Repository().SaveDocument(document, accountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do conta %v", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
