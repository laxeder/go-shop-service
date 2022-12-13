package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

// restaura uma conta de usuário com status deletado
func RestoreUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	document := ctx.Params("document")

	// carrega a conta de usurário com base no documento
	userDatabase, err := user.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário. (%v)", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC097"))
	}

	// verifica o status do usuário
	if userDatabase.Status != user.Disabled {
		log.Error().Msgf("Este usuário já está ativo no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Este usuário já está ativo no sistema."))
	}

	// muda o status do usuário para ativo
	userDatabase.Status = user.Enabled
	userDatabase.UpdatedAt = date.NowUTC()
	userDatabase.Document = document

	// salva as alterações na base de dados
	err = user.Repository().Restore(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC100", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
