package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// assina o Optins da conta
func SignAccountOptins(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// document := ctx.Params("document")

	// // carrega a conta do usuário da base de dados
	// accountData, err := account.Repository().GetDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS101"))
	// }

	// // verifica se o usuário tem uma assinatura á confirmada
	// if accountData.Optins {
	// 	log.Error().Msgf("A assinatura de optins já foi feita para esta conta. (%v)", document)
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS110", "A assinatura de optins já foi feita para esta conta."))
	// }

	// // confirma a assinatura
	// accountData.Optins = true
	// accountData.UpdatedAt = date.NowUTC()

	// // guarda as alterações na conta do usuário na base de dados
	// err = account.Repository().SaveOptins(accountData)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS102"))
	// }

	return response.Ctx(ctx).Result(response.Success(204))
}
