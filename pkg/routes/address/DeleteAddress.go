package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// muda o status do usuário na base de dados
func DeleteAddress(ctx *fiber.Ctx) error {
	var log = logger.New()

	document := ctx.Params("document")

	addressDatabase, err := address.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msg("Os campos enviados estão incorretos.")
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC087"))
	}

	// verifica o status do endereço
	if addressDatabase.Status != address.Enabled {
		log.Error().Msgf("Este endereço já está desativado no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este endereço já está desativado no sistema."))
	}

	addressDatabase.Status = address.Disabled
	addressDatabase.UpdatedAt = date.NowUTC()
	addressDatabase.Document = document

	err = address.Repository().Delete(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
