package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// mosta os dados de uma consta
func ShowAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	uid := ctx.Params("uid")

	accountData, err := account.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC130"))
	}

	return response.Ctx(ctx).Result(response.Success(200, accountData))

}
