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

	uid := ctx.Params("uid")

	accountDatabase, err := account.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar conta. (%v)", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS015"))
	}

	// verifica o status do conta
	if accountDatabase.Status == account.Enabled {
		log.Error().Msgf("Este conta já está ativada no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS016", "Este conta já está ativo no sistema."))
	}

	// muda o status do conta para ativo
	accountDatabase.Uuid = uid
	accountDatabase.Status = account.Enabled
	accountDatabase.UpdatedAt = date.NowUTC()

	// salva as alterações na base de dados
	err = account.Repository().Restore(accountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS017", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
