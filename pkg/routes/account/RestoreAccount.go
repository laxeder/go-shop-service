package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// restaura uma conta de conta com status deletado
func RestoreAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	document := ctx.Params("document")

	// carrega a conta de usurário com base no documento
	accountDatabase, err := account.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar conta. (%v)", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC097"))
	}

	// verifica o status do conta
	if accountDatabase.Status != account.Disabled {
		log.Error().Msgf("Este conta já está ativo no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este conta já está ativo no sistema."))
	}

	// muda o status do conta para ativo
	accountDatabase.Status = account.Enabled
	accountDatabase.Document = document
	accountDatabase.UpdatedAt = date.NowUTC()

	// salva as alterações na base de dados
	err = account.Repository().Restore(accountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC100", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
