package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func ListUsers(ctx *fiber.Ctx) error {

	var log = logger.New()

	users, err := user.Repository().GetList()

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar usuários.")
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS119"))
	}

	return response.Ctx(ctx).Result(response.Success(200, users))

}
