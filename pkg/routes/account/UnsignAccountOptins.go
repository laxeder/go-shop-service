package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// desasina o direito de emitir  CDU
func UnsignAccountOptins(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// document := ctx.Params("document")

	// // carrega uma conta da base de dados
	// accountData, err := account.Repository().GetDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("BLC095"))
	// }

	// // verifica se existe uma assinatura confirmada
	// if !accountData.Optins {
	// 	log.Error().Msgf("A assinatura de optins já está inativa. (%v)", document)
	// 	return response.Ctx(ctx).Result(response.Error(400, "BLC098", "A assinatura de optins já está inativa."))
	// }

	// // desassina o direito de emitor CDu
	// accountData.Optins = false
	// accountData.UpdatedAt = date.NowUTC()

	// // guarda as alterações na conta na base de dados
	// err = account.Repository().SaveOptins(accountData)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("BLC099"))
	// }

	return response.Ctx(ctx).Result(response.Success(204))
}
