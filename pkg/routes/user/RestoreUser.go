package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func RestoreUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	userData, err := user.Repository().GetDataInfo(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS124"))
	}

	if userData == nil {
		log.Error().Msgf("Usuário não encontrado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS184", "Esse usuário não foi encontrado na base de dados."))
	}

	if userData.Status != status.Disabled {
		log.Error().Msgf("Usuário já está ativado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS125", "Este usuário já está ativado no sistema."))
	}

	err = user.Repository().Restore(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar restaurar usuário (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS126"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
