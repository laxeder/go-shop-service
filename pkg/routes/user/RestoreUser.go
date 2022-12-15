package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

// restaura uma conta de usuário com status deletado
func RestoreUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	userDatabase, err := user.Repository().GetUuid(uuid)
	fmt.Print(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário. (%v)", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS097"))
	}

	// verifica o status do usuário
	if userDatabase.Status != user.Disabled {
		log.Error().Msgf("Este usuário já está ativo no sistema. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS060", "Este usuário já está ativo no sistema."))
	}

	// muda o status do usuário para ativo
	userDatabase.Uuid = uuid
	userDatabase.Status = user.Enabled
	userDatabase.UpdatedAt = date.NowUTC()

	// salva as alterações na base de dados
	err = user.Repository().Restore(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS100", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
