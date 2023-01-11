package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListAccounts(ctx *fiber.Ctx) error {

	var log = logger.New()

	accounts, err := account.Repository().GetList()

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar listar contas.")
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS014"))
	}

	return response.Ctx(ctx).Result(response.Success(200, accounts))

}
