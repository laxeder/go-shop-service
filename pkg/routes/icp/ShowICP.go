package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// mostra os dados de uma CDu
func ShowICP(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// document := ctx.Params("document")

	// // se o código da CDu está vazio
	// if str.DocumentValid(document) {
	// 	log.Error().Msg("O documento do usuário não pode ser vazio.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "BLC243", "O documento do usuário não pode ser vazio."))
	// }

	// // carerega a CDU da base de dados
	// icpData, err := icp.Repository().GetByDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error ao tentar encontrar o certificado ICP do usuário" + document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("BLC244"))
	// }

	// // Verifica se a CDo tem uma documento vinculado
	// if len(icpData.Document) == 0 || len(icpData.SerialNumber) == 0 {
	// 	log.Error().Msg("Não existe um certificado ICP registrado para o usuário " + document)
	// 	return response.Ctx(ctx).Result(response.Error(400, "BLC245", "Não existe um certificado ICP registrado para o usuário "+document))
	// }

	// return response.Ctx(ctx).Result(response.Success(200, icpData))
	return response.Ctx(ctx).Result(response.Success(200))

}
