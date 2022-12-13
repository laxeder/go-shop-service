package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func UpdateUserDocument(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// iniicia a struct do usuário com password
	userDocument := &user.UserDocument{}

	// converte o json para struct
	err := json.Unmarshal(body, userDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "O formado dos dados envidados está incorreto."))
	}

	// carrega os dados do usuário da base de dados
	userDatabase, err := user.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Documento não encontrado %v.", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC093"))
	}

	if userDocument.OldDocument != document {
		log.Error().Err(err).Msgf("Documento antigo errado %v.", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "Documento antigo está incorreto."))
	}

	// define o novo documento
	userDatabase.SetDocument(userDocument.Document)
	userDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do usuário
	err = user.Repository().SaveDocument(document, userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do usuário %v", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
