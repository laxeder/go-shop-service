package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// restaura um endereço com status deletado
func RestoreAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	document := ctx.Params("document")

	// carrega o endereço com base no documento
	addressDatabase, err := address.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar endereço. (%v)", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC097"))
	}

	// verifica o status do endereço
	if addressDatabase.Status != address.Disabled {
		log.Error().Msgf("Este endereço já está ativo no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este endereço já está ativo no sistema."))
	}

	// muda o status do endereço para ativo
	addressDatabase.Status = address.Enabled
	addressDatabase.Document = document
	addressDatabase.UpdatedAt = date.NowUTC()

	// salva as alterações na base de dados
	err = address.Repository().Restore(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC100", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
