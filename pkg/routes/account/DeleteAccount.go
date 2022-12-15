package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// muda o status do conta na base de dados
func DeleteAccount(ctx *fiber.Ctx) error {
	var log = logger.New()

	document := ctx.Params("document")

	accountDatabase, err := account.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC087"))
	}

	// verifica o status da conta
	if accountDatabase.Status == account.Disabled {
		log.Error().Msgf("Esta conta já está desativado no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Esta conta já está desativado no sistema."))
	}

	accountDatabase.Status = account.Disabled
	accountDatabase.UpdatedAt = date.NowUTC()
	accountDatabase.Document = document

	err = account.Repository().Delete(accountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
