package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateAddressDocument(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// iniicia a struct do endereço com password
	addressDocument := &address.AddressDocument{}

	// converte o json para struct
	err := json.Unmarshal(body, addressDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "O formado dos dados envidados está incorreto."))
	}

	// carrega os dados do endereço da base de dados
	addressDatabase, err := address.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Documento não encontrado %v.", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC093"))
	}

	if addressDocument.OldDocument != document {
		log.Error().Err(err).Msgf("Documento antigo errado %v.", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "Documento antigo está incorreto."))
	}

	// define o novo documento
	addressDatabase.SetDocument(addressDocument.Document)
	addressDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do endereço
	err = address.Repository().SaveDocument(document, addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do endereço %v", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
