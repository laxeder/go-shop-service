package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func ShowUser(ctx *fiber.Ctx) error {
	var log = logger.New()

	uuid := ctx.Params("uuid")

	userData, err := user.Repository().Get(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS127"))
	}

	if userData == nil {
		log.Error().Msgf("Usuário não encontrado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS186", "Esse usuário não foi encontrado na base de dados."))
	}

	return response.Ctx(ctx).Result(response.Success(200, userData))

}
