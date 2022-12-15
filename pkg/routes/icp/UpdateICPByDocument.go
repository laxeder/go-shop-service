package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateICPByDocument(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// body := ctx.Body()
	// document := ctx.Params("document")

	// // converte json para struct
	// icpBody, err := icp.New(body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS247", "O formado dos dados envidados está incorreto."))
	// }

	// // verifica e compara o documento recebido
	// if icpBody.Document != "" && str.DocumentClean(document) != str.DocumentClean(icpBody.Document) {
	// 	log.Error().Msg("Não é possível atualizar o certificado com  o documento " + document)
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS248", "Não é possível atualizar o certificado com  o documento "+document))
	// }

	// // ! ###########################################################################################################################
	// // ! VALIDA OS DADOS DE ENTRADA
	// // ! ###########################################################################################################################

	// if icpBody.KeyPublic != "" {
	// 	checkKeyPublic := icpBody.KeyPublicValid()
	// 	if checkKeyPublic.Status != 200 {
	// 		return response.Ctx(ctx).Result(checkKeyPublic)
	// 	}
	// }

	// if icpBody.Name != "" {
	// 	checkName := icpBody.NameValid()
	// 	if checkName.Status != 200 {
	// 		return response.Ctx(ctx).Result(checkName)
	// 	}
	// }

	// if icpBody.Email != "" {
	// 	checkEmail := icpBody.EmailValid()
	// 	if checkEmail.Status != 200 {
	// 		return response.Ctx(ctx).Result(checkEmail)
	// 	}
	// }

	// if icpBody.Validate != "" {
	// 	checkValidate := icpBody.ValidateValid()
	// 	if checkValidate.Status != 200 {
	// 		return response.Ctx(ctx).Result(checkValidate)
	// 	}
	// }

	// // carrega o certificado da base de dados
	// icpData, err := icp.Repository().GetByDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Erro ao tentar encontrar o certificado.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS249", "Erro ao tentar encontrar o certificado."))
	// }

	// // injecta dos dados novos no lugar dos dados trazidos da base de dados
	// icpData.Inject(icpBody)
	// icpData.Validate = date.BRToUTC(icpData.Validate)
	// icpData.UpdatedAt = date.NowUTC()

	// // guarda as alterações do ceritificado na base de dados
	// err = icp.Repository().Update(icpData)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS250"))
	// }

	return response.Ctx(ctx).Result(response.Success(204))
}
