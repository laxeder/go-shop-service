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

	document := ctx.Params("document")

	// carrega a conta do usuário com vase no documento passdo
	accountData, err := account.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC130"))
	}

	return response.Ctx(ctx).Result(response.Success(200, accountData))

}
