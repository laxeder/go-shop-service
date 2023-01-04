package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func DeleteUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")

	userInfo, err := user.Repository().GetDataInfo(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS116"))
	}

	if userInfo == nil {
		log.Error().Msgf("Usuário não encontrado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS185", "Esse usuário não foi encontrado na base de dados."))
	}

	if userInfo.Status != status.Enabled {
		log.Error().Msgf("Usuário já está desativado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS117", "Este usuário já está desativado no sistema."))
	}

	err = user.Repository().Delete(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar deletar usuário (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS118"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
